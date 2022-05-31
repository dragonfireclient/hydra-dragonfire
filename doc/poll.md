# Polling API
Source code: [poll.go](../poll.go)

`poll` waits for and returns the next event from one or more clients, or `nil` if none of the clients passed to it are active (`connected` state).
Optionally, a timeout can be passed to poll; if no other event occurs until the timeout elapses, a timeout event is returned.

## Events

An event is a table that contains a string `type`. Depending on the type, it may have different other fields.

- `type = "interrupt"`: Fired globally when the program was interrupted using a signal.

- `type = "timeout"`: Fired when the timeout elapses.

- `type = "pkt"`: Fired when a packet was received. See [pkts.md](pkts.md)

- `type = "disconnect"`: Fired when a client connection closed. Has a `client` field. 

- `type = "error"`: Fired when an error occurs during deserialization of a packet. Has a `client` field. Stores the error message in an `error` field.
