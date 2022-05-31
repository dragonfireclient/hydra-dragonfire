# Map Component
Source code: [map.go](../map.go)

Map handles the `blk_data` and `node_metas_changed` packets.
Map may send `got_blks`, `deleted_blks` packets.

## Functions

`self:clear()`: Forget all blocks.

`self:block(blkpos)`: Return the `map_blk` at `blkpos` as found in the `blk_data` packet (See [client_pkts.md](client_pkts.md)). `nil` if block is not present.

`self:node(pos)`: Return a node in the form `{ param0 = 126, param1 = 0, param2 = 0, meta = { ... } }`. The meta field is a `node_meta` as found in the `blk_data` packet. `nil` if node is not present.
