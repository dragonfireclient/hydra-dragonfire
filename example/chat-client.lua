#!/usr/bin/env hydra-dragonfire
local escapes = require("escapes")
local client = require("client")()

client:subscribe("chat_msg")
client:connect()

while not hydra.canceled() do
	local pkt, interrupt = client:poll(1)

	if pkt then
		print(escapes.strip_all(pkt.text))
	elseif interrupt then
		client:send("chat_msg", {msg = "test"})
	else
		print("disconnected")
		break
	end
end

client:close()
