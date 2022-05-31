# Hydra namespace
Source code: [hydra.go](../hydra.go)

The `hydra` table contains functions necessary to handle connections.

## Constants

- `hydra.BS`: (`= 10.0`) Size of node in floating-point units. Many floating point positions contained in packets received need to be divided by this constant to obtain a position in node space. Likewise, many positions contained in packets sent need to be multiplied by this constant.
- `hydra.serialize_ver`: Supported serialization version.
- `hydra.proto_ver`: Supported protocol version.

## Functions

- `hydra.client(address)`: Returns a new client. Address must be a string. For client functions, see [client.md](client.md).
- `hydra.dtime()`: Utility function that turns the elapsed time in seconds (floating point) since it was last called (or since program start).
- `hydra.poll(clients, [timeout])`: Polls events from all clients in `clients` (table). For behavior and return value, see [poll.md](poll.md).
- `hydra.close(clients)`: Closes all clients in `clients` (table) that are currently connected. See `client:close()` in [client.md](client.md) for more info.
