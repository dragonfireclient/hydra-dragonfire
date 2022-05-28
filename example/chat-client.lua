#!/usr/bin/env hydra-dragonfire
local escapes = require("escapes")
local address, name, password = unpack(arg)
local client = hydra.client(address)

client:enable("auth")
client.auth:username(name)
client.auth:password(password or "")

client:subscribe("chat_msg")
client:connect()

while not hydra.canceled() do
	local pkt, interrupt = client:poll()

	if pkt then
		print(escapes.strip_all(pkt.text))
	elseif not interrupt then
		print("disconnected")
		break
	end
end

client:disconnect()
