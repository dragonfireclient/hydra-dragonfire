# Packets Component
Source code: [pkts.go](../pkts.go)

The packets component allows you to handle packets yourself. It fires events in the form of `{ type = "pkt", client = ..., pkt_type = "...", pkt_data = { ... } }``` when subscribed packets are received.
For available packets, see [client_pkts.md](client_pkts.md). By default, no packets are packets subscribed.

## Wildcard mode

If wildcard is enabled, events for all packets are fired, even ones that are not subscribed. It is not recommended to use this without a reason since converting packets to Lua costs performance and creates and overhead due to poll returning more often. `wildcard` is unnecessary if only certain packets are handled anyway, but it is useful for traffic inspection and debugging.

## Functions

- `self:subscribe(pkt1, [pkt2, ...])`: Subscribes to all packet types passed as arguments (strings).

- `self:unsubscribe(pkt1, [pkt2, ...])`: Unsubscribes from all packet passed as arguments (strings).

- `self:wildcard(wildcard)`: Sets wildcard mode to `wildcard` (boolean).

