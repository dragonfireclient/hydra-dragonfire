--[[ builtin/client.lua ]]--

function package.loaded.client()
	local address, name, password = unpack(arg)
	local client = hydra.client(address)

	client:enable("auth")
	client.auth:username(name)
	client.auth:password(password or "")

	return client
end
