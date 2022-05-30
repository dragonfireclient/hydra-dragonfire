# Polling API
Source code: [poll.go](../poll.go)

**TL;DR**: poll is complex and has many different cases, but in general, it returns the received packet and the associated client; if one of the clients closes, a nil packet is returned once. client may also be nil in some cases so watch out for that.

Together with sending, polling is the core function of hydra. It is used to receive packets from a packet queue.

For each client, only packets that the client has subscribed to are inserted into that queue, unless wildcard is enabled.

Packet receival from network happens asynchronously. When a packet is received and has been processed by components, it is enqueued for polling if the client is subscribed to it. **Because of the poll queue, packets may be returned by poll that the client was subscribed to in the past but unsubscribed recently.** Since the queue has a limited capacity of 1024 packets (this may change in the future), it is your responsibility to actually poll in a frequency suitable to keep up with the amount of packets you expect based on what you are subscribed to. If the queue is full, the thread responsible for receival will block.

Clients that are not in `connected` state are ignored by poll.

Poll blocks until one of these conditions is met (in this order). The return value depends on which condition is met:

1. No clients are available when the function is called. This happens if either no clients were passed to poll or none of them is connected.

2. One of the clients closes. In this case, the client that closed is set to `disconnected` state. The close may happen before or during the call to poll, but it has effect only once.

3. A packet is in queue for one of the clients (Main case).

4. An interrupt signal is received during polling (See `hydra.canceled`).

5. The configured timeout elapses.

## Different versions

There is two different versions of poll: `client:poll` for polling a single client and `hydra.poll` for polling multiple clients.
They are mostly equivalent but differ in return values and arguments:

- `client:poll([timeout])` polls from the client `client` and returns `pkt, interrupted`

- `hydra.poll(clients, [timeout])` takes table of clients as argument and returns `pkt, client, interrupted`

## Arguments and return values

The timeout argument is an optional floating point number holding the timeout in seconds, if `nil`, poll will block until one of the conditions 1.-4. are met. Timeout may be `0`, in this case poll returns immediately even if none of the other conditions are met immediately.

Return values for different cases:

1. If no clients are available, `nil, nil, false` (or `nil, false` respectively) is returned.

2. If a client closes, `nil, client, false` (or `nil, false` respectively) is returned.

3. If a packet is available, poll returns `pkt, client, false` (or `pkt, false` respectively). `pkt` is a table containing the received packet (see [client_pkts.md](client_pkts.md)) and `client` is the client reference that has received the packet.

4. If the program is interrupted, poll returns `nil, nil, true` (or `nil, true` respectively).

5. If the timeout elapses, poll returns `nil, nil, true` (or `nil, true` respectively).
