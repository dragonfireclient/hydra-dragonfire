#!/usr/bin/env hydra-dragonfire
local client = require("client")()

client:enable("pkts", "map")
client.pkts:subscribe("move_player")

client:connect()

local pos

while true do
	local evt = client:poll(1)

	if not evt or evt.type == "disconnect" or evt.type == "interrupt" then
		break
	elseif evt.type == "pkt" then
		pos = (evt.pkt_data.pos / hydra.BS + vec3(0, -1, 0)):round()
	elseif evt.type == "timeout" and pos then
		local node = client.map:node(pos)
		print(pos, node and node.param0)
	end
end

client:close()
