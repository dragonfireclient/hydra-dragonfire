# Hydra
<img src="https://cdn8.picryl.com/photo/2016/05/14/hydra-from-bl-royal-12-c-xix-f-13-b31e4a-1024.jpg" width="500" />

Lua bindings for client side minetest protocol - written in Go, using [anon5's mt package](https://github.com/anon55555/mt).
Capable of deserializing ToClt packets and serializing ToSrv packets.

Main use case are bot clients. Multiple clients can easily be spawend and polled from one script - hence the name "Hydra".
Hydra may also be useful for integration testing or network debugging.

Hydra is WIP: there are bugs, API may change any time, doc is incomplete, some packets are unimplemented and many components are yet to be added. However, hydra can already be used and big parts of it's main functionality are implemented.

# Installation
Go 1.18 is required.
`go install github.com/dragonfireclient/hydra-dragonfire@latest`

# Invocation
Due to limitations of Go, hydra unfortunately cannot be `require()`'d from a Lua script. Instead, the hydra binary has to be invoked with a script as argument:
`hydra-dragonfire file.lua <args>`. Any additional arguments `<args>` are provided to the script.

# Architecture Overview
By default, hydra will only take care of connection and packet serialization, no state management takes place.
Hydra is a library, not a framework: it does not organize your code and there is no module system.

Multiple clients can be created independently by calling the `hydra.client` function.
`poll` can be used on one or multiple clients to receive packets and `send` can be used to send packets.
Only selected packets will be returned by `poll`, to avoid unnecessary conversion of packets to Lua.
Poll will return early if the script is interrupted by a signal, one of the selected clients is disconnected or the configured timeout elapses.

Additionally, different native components can be enabled per-client to manage state.
Currently only the `auth` component is available, but components like `map`, `objs`, `inv`, `pos`, `playerlist` etc. will be added in the future.
Components handle packets asynchronously, they will process them even if poll is not called.

# Further Documentation
[API documentation](doc/api.md)
