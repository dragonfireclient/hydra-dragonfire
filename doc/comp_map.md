# Map Component
Source code: [comp_map.go](../comp_map.go)

The Map component stores a reference to a `hydra.map` (See [map.md](map.md)).

Initially, an empty map is created. You can replace this by a map reference obtained from `hydra.map` however: this way, multiple clients can share a map and explore different areas of it.

Handles the `blk_data` and `node_metas_changed` packets.
May send `got_blks` packets.

## Functions

`self:set(mapref)`: Data will be stored in `mapref` in the future.

`self:get()`: Returns the current `mapref`.

