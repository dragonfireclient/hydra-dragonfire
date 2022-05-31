#!/usr/bin/env hydra-dragonfire
local escapes = require("escapes")
local base64 = require("base64")
local client = require("client")()

client:wildcard(true)
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
		print(val._type or "")

		local idt = (indent or "") .. "  "
		for k, v in pairs(val) do
			if k ~= "_type" then
				io.write(idt .. k .. " ")
				dump(v, idt)
			end
		end
	end
end

while not hydra.canceled() do
	local pkt, interrupt = client:poll()

	if pkt then
		if pkt._type == "srp_bytes_salt_b" then
			pkt.b = base64.encode(pkt.b)
			pkt.salt = base64.encode(pkt.salt)
		end

		if pkt._type == "chat_msg" then
			pkt.text = escapes.strip_all(pkt.text)
		end

		if pkt._type == "blk_data" then
			pkt.blk.param0 = {}
			pkt.blk.param1 = {}
			pkt.blk.param2 = {}
		end

		dump(pkt)
	elseif not interrupt then
		print("disconnected")
		break
	end
end

client:close()
