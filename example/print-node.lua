#!/usr/bin/env hydra-dragonfire
local client = require("client")()
client:enable("map")

client:subscribe("move_player")
client:connect()

local pos

while not hydra.canceled() do
	local pkt, interrupted = client:poll(1)

	if pkt then
		pos = (pkt.pos / hydra.BS + vec3(0, -1, 0)):round()
	elseif not interrupted then
		break
	elseif pos then
		local node = client.map:node(pos)
		print(pos, node and node.param0)
	end
end

client:close()
