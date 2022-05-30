#!/usr/bin/env hydra-dragonfire
local client = require("client")()

client:wildcard(true)
client:connect()

while not hydra.canceled() do
	local pkt, interrupt = client:poll()

	if pkt then
		print(pkt._type)
		for k, v in pairs(pkt) do
			if k ~= "_type" then
				print("", k, v)
			end
		end
	elseif not interrupt then
		print("disconnected")
		break
	end
end

client:close()
