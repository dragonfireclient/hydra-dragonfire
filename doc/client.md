# Client Reference
Source code: [client.go](../client.go)

A client represents a connection state. A client may not yet be connected, or it may be already disconnected.
After being disconnect, a client cannot be reconnected.

## Functions

- `self:address()`: Returns address passed to `hydra.client` upon creation as a string.
- `self:state()`: Returns current connection state as a string ("new", "connected", "disconnected")
- `self:connect()`: Connects to server. Throws an error if the client is not in "new" state OR address resolution / dial fails (Note: If required, you can use `pcall` to catch and handle errors instead of crashing the script). Connection failure (= host found, but no minetest server running on port) is equivalent to an immediate disconnect and does not cause an error to be thrown.
- `self:poll([timeout])`: Polls events from client. See [poll.md](poll.md) for behavior and return values.
- `self:close()`: Closes the network connection if in `connected` state. The client remains in `connected` state until passed to poll.
- `self:enable(component)`: Enables the component with the name `component` (string), if not already enabled. By default, no components are enabled. See Components section.
- `self:send(pkt_type, pkt_data, [ack])`: Sends a packet to server. Throws an error if the client is not connected. `pkt_type` is the type of the packet as string. `pkt_data` is a table containing packet parameters. Some packets don't have parameters (e.g. `respawn`) - in this case, `pkt_data` can be omitted. See [server_pkts.md](server_pkts.md) for available packets. If `ack` is true, this function will block until acknowledgement from server is received.

## Components

Enabled components can be accessed by using `self.<component name>`.

- `self.pkt`: Allows you to handle selected packets yourself. Most scripts use this. See [comp_pkts.md](comp_pkts.md).
- `self.auth`: Handles authentication. Recommended for the vast majority of scripts. See [comp_auth.md](comp_auth.md).
- `self.map`: Stores MapBlocks received from server. See [comp_map.md](comp_map.md).
