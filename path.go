package main

import (
	"github.com/beefsack/go-astar"
	"github.com/dragonfireclient/hydra-dragonfire/convert"
	"github.com/dragonfireclient/mt"
	"github.com/yuin/gopher-lua"
	"math"
	"sync"
)

type PathNode struct {
	pos   [3]int16
	blk   *MapBlk
	edges map[*PathNode]struct{}
}

const pathMaxTp float64 = 4.317 * 10.0 * 0.5
const pathMaxTpSq float64 = pathMaxTp * pathMaxTp

var pathDirs = [6][3]int16{
	[3]int16{+1, 0, 0},
	[3]int16{-1, 0, 0},
	[3]int16{0, +1, 0},
	[3]int16{0, -1, 0},
	[3]int16{0, 0, +1},
	[3]int16{0, 0, -1},
}

func pathAddPos(a, b [3]int16) [3]int16 {
	for i, v := range a {
		b[i] += v
	}
	return b
}

func pathScalePos(v [3]int16, s int16) (r [3]int16) {
	for i, x := range v {
		r[i] = x * s
	}
	return
}

func pathDistSq(a, b [3]int16) float64 {
	distSq := 0.0

	for i, v := range a {
		if i != 1 {
			abs := math.Abs(float64(v - b[i]))
			if abs > 0 {
				abs -= 1
			}
			distSq += abs
		}
	}

	return distSq
}

func pathAddEdge(a, b *PathNode) {
	a.edges[b] = struct{}{}
	b.edges[a] = struct{}{}
}

func pathAddNode(blk *MapBlk, pos [3]int16) (node *PathNode, ok bool) {
	node, ok = blk.pathNodes[pos]
	if ok {
		return
	}

	node = &PathNode{}
	node.pos = pos
	node.blk = blk
	node.edges = map[*PathNode]struct{}{}

	blk.pathNodes[pos] = node
	return
}

func pathRemoveEdge(from, to *PathNode) {
	delete(from.edges, to)

	// garbage collect
	if len(from.edges) == 0 {
		pathRemoveNode(from)
	}
}

func pathRemoveNode(node *PathNode) {
	for nbr := range node.edges {
		pathRemoveEdge(nbr, node)
	}

	if node.blk != nil {
		delete(node.blk.pathNodes, node.pos)
	}
}

func pathCheckAddEdge(blk1, blk2 *MapBlk, pos1, pos2 [3]int16, mu *sync.Mutex) bool {
	if pathDistSq(pos1, pos2) > pathMaxTpSq {
		return false
	}

	mu.Lock()
	defer mu.Unlock()

	node1, _ := pathAddNode(blk1, pos1)
	node2, _ := pathAddNode(blk2, pos2)

	pathAddEdge(node1, node2)
	return true
}

func pathAddBlock(mp *Map, blk1 *MapBlk, blkpos1 [3]int16) {
	blk1.pathNodes = map[[3]int16]*PathNode{}
	nodpos1 := pathScalePos(blkpos1, 16)

	// sync stuff
	var chans []chan [3]int16
	var wg sync.WaitGroup
	var mu sync.Mutex
	var done bool

	for _, dir := range pathDirs {
		blkpos2 := pathAddPos(blkpos1, dir)
		nodpos2 := pathScalePos(blkpos2, 16)

		blk2, ok := mp.blocks[blkpos2]
		if !ok {
			continue
		}

		ch := make(chan [3]int16, 4096)
		chans = append(chans, ch)
		wg.Add(1)

		go func() {
			defer wg.Done()

			var positions [][3]int16
			for x := uint16(0); x < 16; x++ {
				for z := uint16(0); z < 16; z++ {
					for y := uint16(0); y < 16; y++ {
						if blk2.data.Param0[x|(y<<4)|(z<<8)] != mt.Air {
							continue
						}

						pos2 := pathAddPos(nodpos2, [3]int16{int16(x), int16(y), int16(z)})

						for _, pos1 := range positions {
							if pathCheckAddEdge(blk1, blk2, pos1, pos2, &mu) {
								return
							}
						}

						for ch != nil {
							pos1, ok := <-ch
							if ok {
								if pathCheckAddEdge(blk1, blk2, pos1, pos2, &mu) {
									return
								} else {
									positions = append(positions, pos1)
								}
							} else {
								ch = nil
								if len(positions) == 0 {
									return
								}
							}
						}
					}
				}
			}
		}()
	}

	go func() {
		for _, ch := range chans {
			defer close(ch)
		}

		for x := uint16(0); x < 16; x++ {
			for z := uint16(0); z < 16; z++ {
				for y := uint16(0); y < 16; y++ {
					if done {
						return
					}

					if blk1.data.Param0[x|(y<<4)|(z<<8)] != mt.Air {
						continue
					}

					for _, ch := range chans {
						ch <- pathAddPos(nodpos1, [3]int16{int16(x), int16(y), int16(z)})
					}
					break
				}
			}
		}
	}()

	wg.Wait()
	done = true
}

func pathRemoveBlock(blk *MapBlk) {
	for _, node := range blk.pathNodes {
		node.blk = nil
		pathRemoveNode(node)
	}
}

func (node *PathNode) PathNeighbors() (edges []astar.Pather) {
	for node := range node.edges {
		edges = append(edges, node)
	}
	for _, node := range node.blk.pathNodes {
		edges = append(edges, node)
	}
	return
}

func (node *PathNode) PathNeighborCost(to astar.Pather) float64 {
	return node.PathEstimatedCost(to)
}

func (node *PathNode) PathEstimatedCost(to astar.Pather) float64 {
	return pathDistSq(node.pos, to.(*PathNode).pos)
}

func pathFind(mp *Map, pos [2][3]int16, l *lua.LState) lua.LValue {
	var abs [2]struct {
		blk  *MapBlk
		node *PathNode
		ex   bool
	}

	for i := range abs {
		blkpos, _ := mt.Pos2Blkpos(pos[i])
		blk, ok := mp.blocks[blkpos]
		if !ok {
			return lua.LNil
		}

		abs[i].node, abs[i].ex = pathAddNode(blk, pos[i])
	}

	// reverse dst and src due to bug in astar package
	path, _, found := astar.Path(abs[1].node, abs[0].node)
	if !found {
		return lua.LNil
	}

	for i := range abs {
		if !abs[i].ex {
			pathRemoveNode(abs[i].node)
		}
	}

	tbl := l.NewTable()
	for i, pather := range path {
		pos := pather.(*PathNode).pos
		lpos := [3]lua.LNumber{lua.LNumber(pos[0]), lua.LNumber(pos[1]), lua.LNumber(pos[2])}

		l.RawSetInt(tbl, i+1, convert.PushVec3(l, lpos))
	}
	return tbl
}
