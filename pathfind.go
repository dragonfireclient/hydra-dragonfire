package main

import (
	"github.com/anon55555/mt"
	"math"
	"sync"
)

type PathfindEdge struct {
	src, dst [3]int16
	weight   float64
}

const pfMaxtp float64 = 4.317 * 10.0 * 0.5
const pfMaxtpSq float64 = pfMaxtp * pfMaxtp

var pfDirs = [6][3]int16{
	[3]int16{+1, 0, 0},
	[3]int16{-1, 0, 0},
	[3]int16{0, +1, 0},
	[3]int16{0, -1, 0},
	[3]int16{0, 0, +1},
	[3]int16{0, 0, -1},
}

func pfRidx(idx int) int {
	if idx%2 == 0 {
		return idx + 1
	} else {
		return idx - 1
	}
}

func pfCenterFindAir(blk *MapBlk, pos [3]int16, chans [6]chan [3]int16, done *bool) {
	for _, ch := range chans {
		if ch != nil {
			defer close(ch)
		}
	}

	for x := uint16(0); x < 16; x++ {
		for z := uint16(0); z < 16; z++ {
			for y := uint16(0); y < 16; y++ {
				if *done {
					return
				}

				if blk.data.Param0[x|(y<<4)|(z<<8)] == mt.Air {
					for _, ch := range chans {
						if ch != nil {
							ch <- [3]int16{int16(x) + pos[0], int16(y) + pos[1], int16(z) + pos[2]}
						}
					}
					break
				}
			}
		}
	}
}

func pfMakeEdge(src [3]int16, dst [3]int16, vertical bool, edge **PathfindEdge) bool {
	var distSq float64

	for i, v := range dst {
		if vertical == (i == 1) {
			abs := math.Abs(float64(v - src[i]))
			if abs > 0 {
				abs -= 1
			}
			distSq += math.Pow(abs, 2)
		}
	}

	if vertical || distSq <= pfMaxtpSq {
		*edge = &PathfindEdge{
			src:    src,
			dst:    dst,
			weight: math.Sqrt(distSq),
		}

		return true
	}

	return false
}

func pfNeighFindAir(blk *MapBlk, pos [3]int16, ch chan [3]int16, wg *sync.WaitGroup, vertical bool, edge **PathfindEdge) {
	defer wg.Done()

	var prev [][3]int16

	for x := uint16(0); x < 16; x++ {
		for z := uint16(0); z < 16; z++ {
			for y := uint16(0); y < 16; y++ {
				if blk.data.Param0[x|(y<<4)|(z<<8)] == mt.Air {
					dst := [3]int16{int16(x) + pos[0], int16(y) + pos[1], int16(z) + pos[2]}

					for _, src := range prev {
						if pfMakeEdge(dst, src, vertical, edge) {
							return
						}
					}

					for ch != nil {
						src, ok := <-ch
						if ok {
							if pfMakeEdge(dst, src, vertical, edge) {
								return
							} else {
								prev = append(prev, src)
							}
						} else {
							ch = nil
							if len(prev) == 0 {
								return
							}
						}
					}
				}
			}
		}
	}
}

func pfPreprocess(mp *Map, blkpos [3]int16, blk *MapBlk) {
	println("preprocess")

	var chans [6]chan [3]int16
	var blks [6]*MapBlk
	var wg sync.WaitGroup

	var pos [3]int16
	for k, v := range blkpos {
		pos[k] = v * 16
	}

	for i := range chans {
		npos := pos
		nblkpos := blkpos
		for j, v := range pfDirs[i] {
			npos[j] += v * 16
			nblkpos[j] += v
		}

		if nblk, ok := mp.blocks[nblkpos]; ok {
			blks[i] = nblk
			chans[i] = make(chan [3]int16, 4096)
			wg.Add(1)
			go pfNeighFindAir(blk, npos, chans[i], &wg, i == 2 || i == 3, &blk.edges[i])
		}
	}

	var done bool
	go pfCenterFindAir(blk, pos, chans, &done)
	wg.Wait()
	done = true

	for i, nblk := range blks {
		if nblk != nil {
			edge := blk.edges[i]
			ri := pfRidx(i)

			if edge == nil {
				nblk.edges[ri] = nil
			} else {
				nblk.edges[ri] = &PathfindEdge{
					src:    edge.dst,
					dst:    edge.src,
					weight: edge.weight,
				}
			}
		}
	}

	println("finish preprocess")
}
