#!/usr/bin/env hydra-dragonfire
local escapes = require("escapes")
local client = require("client")()

client:enable("pkts")
client.pkts:subscribe("chat_msg")

client:connect()

while true do
	local evt = client:poll(1)

	if not evt then
		break
	end

	if not evt or evt.type == "interrupt" or evt.type == "disconnect" then
		break
	elseif evt.type == "pkt" then
		print(escapes.strip_all(evt.pkt_data.text))
	elseif evt.type == "timeout" then
		client:send("chat_msg", {msg = "test"})
	end
end

client:close()
