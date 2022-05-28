local function snext(t, state)
	local key

	if state == nil then
		t.__sorted = {}
		for k in pairs(t) do
			if k ~= "__sorted" then
				table.insert(t.__sorted, k)
			end
		end
		table.sort(t.__sorted)
		
		key = t.__sorted[1]
	else
		for i, v in ipairs(t.__sorted) do
			if v == state then
				key = t.__sorted[i + 1]
				break
			end
		end
	end

	if key then
		return key, t[key]
	end

	t.__sorted = nil
end

function spairs(t)
	return snext, t, nil
end

local function parse_pair(pair, value_first)
	if pair:sub(1, 1) == "#" then
		return
	end

	local idx = pair:find(" ")

	if idx then
		local first, second = pair:sub(1, idx - 1), pair:sub(idx + 1)

		if value_first and first:sub(1, 1) ~= "[" then
			return second, first
		else
			return first, second
		end
	else
		return pair
	end
end

function parse_spec(name, value_first)
	local f = io.open("../spec/" .. name, "r")
	local spec = {}
	local top

	for l in f:lines() do
		if l:sub(1, 1) == "\t" then
			local key, val = parse_pair(l:sub(2), value_first)

			if val then
				top[key] = val
			elseif key then
				table.insert(top, key)
			end
		else
			local key, val = parse_pair(l, value_first)

			if val then
				spec[key] = val
			elseif key then
				top = {}
				spec[key] = top
			end
		end
	end

	f:close()
	return spec
end

local casemap = parse_spec("casemap")

function camel_case(snake)
	if casemap[snake] then
		return casemap[snake]
	end

	local camel = ""

	while #snake > 0 do
		local idx = snake:find("_") or #snake + 1

		camel = camel
			.. snake:sub(1, 1):upper()
			.. snake:sub(2, idx - 1)

		snake = snake:sub(idx + 1)
	end

	return camel
end

