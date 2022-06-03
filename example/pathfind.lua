#!/usr/bin/env hydra-dragonfire
local client = require("client")()

client:enable("map")
client.map:set(hydra.map(true))

client:enable("pkts")
client.pkts:subscribe("chat_msg")

client:connect()

while true do
	local evt = client:poll()

	if not evt or evt.type == "disconnect" or evt.type == "interrupt" then
		break
	elseif evt.type == "pkt" then
		print("chatmsg")
	end
end

client:close()
