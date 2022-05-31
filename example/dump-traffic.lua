#!/usr/bin/env hydra-dragonfire
local escapes = require("escapes")
local base64 = require("base64")
local client = require("client")()

client:enable("pkts")
client.pkts:wildcard(true)

client:connect()

local function dump(val, indent)
	local t = type(val)
	local mt = getmetatable(val)

	if t ~= "table" or mt and mt.__tostring then
		if t == "string" then
			val = val:gsub("\n", "\\n")
		end
		print(val)
	else
		print()

		local idt = (indent or "") .. "  "
		for k, v in pairs(val) do
			io.write(idt .. k .. " ")
			dump(v, idt)
		end
	end
end

while true do
	local evt = client:poll()

	if not evt or evt.type == "disconnect" or evt.type == "interrupt" then
		break
	elseif evt.type == "error" then
		print(evt.error)
	elseif evt.type == "pkt" then
		local type, data = evt.pkt_type, evt.pkt_data

		if type == "srp_bytes_salt_b" then
			data.b = base64.encode(data.b)
			data.salt = base64.encode(data.salt)
		end

		if type == "chat_msg" then
			data.text = escapes.strip_all(data.text)
		end

		if type == "blk_data" then
			data.blk.param0 = {}
			data.blk.param1 = {}
			data.blk.param2 = {}
		end

		io.write(type)
		dump(data)
	end
end

client:close()
