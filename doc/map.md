# Map Reference
Source code: [map.go](../map.go)

A map stores mapblocks received from server. To be of any use, a map needs to be added to one or more clients. See [comp_map.md](comp_map.md) and [client.md](client.md).

## Functions

`self:block(blkpos)`: Return the `map_blk` at `blkpos` as found in the `blk_data` packet (See [client_pkts.md](client_pkts.md)). `nil` if block is not present.

`self:node(pos)`: Return a node in the form `{ param0 = 126, param1 = 0, param2 = 0, meta = { ... } }`. The meta field is a `node_meta` as found in the `blk_data` packet. `nil` if node is not present.
