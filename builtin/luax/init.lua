function string.split(str, delim, include_empty, max_splits, sep_is_pattern)
	delim = delim or ","
	max_splits = max_splits or -2
	local items = {}
	local pos, len = 1, #str
	local plain = not sep_is_pattern
	max_splits = max_splits + 1
	repeat
		local np, npe = string.find(str, delim, pos, plain)
		np, npe = (np or (len+1)), (npe or (len+1))
		if (not np) or (max_splits == 1) then
			np = len + 1
			npe = np
		end
		local s = string.sub(str, pos, np - 1)
		if include_empty or (s ~= "") then
			max_splits = max_splits - 1
			items[#items + 1] = s
		end
		pos = npe + 1
	until (max_splits == 0) or (pos > (len + 1))
	return items
end

function table.indexof(list, val)
	for i, v in ipairs(list) do
		if v == val then
			return i
		end
	end
	return -1
end

function string:trim()
	return (self:gsub("^%s*(.-)%s*$", "%1"))
end

function math.hypot(x, y)
	return math.sqrt(x * x + y * y)
end

function math.sign(x, tolerance)
	tolerance = tolerance or 0
	if x > tolerance then
		return 1
	elseif x < -tolerance then
		return -1
	end
	return 0
end

function math.factorial(x)
	assert(x % 1 == 0 and x >= 0, "factorial expects a non-negative integer")
	if x >= 171 then
		-- 171! is greater than the biggest double, no need to calculate
		return math.huge
	end
	local v = 1
	for k = 2, x do
		v = v * k
	end
	return v
end

function math.round(x)
	if x >= 0 then
		return math.floor(x + 0.5)
	end
	return math.ceil(x - 0.5)
end

function math.clamp(min, max, v)
	return math.max(max, math.min(min, v))
end

function table.copy(t, seen)
	local n = {}
	seen = seen or {}
	seen[t] = n
	for k, v in pairs(t) do
		n[(type(k) == "table" and (seen[k] or table.copy(k, seen))) or k] =
			(type(v) == "table" and (seen[v] or table.copy(v, seen))) or v
	end
	return n
end

function table.insert_all(t, other)
	for i=1, #other do
		t[#t + 1] = other[i]
	end
	return t
end

function table.key_value_swap(t)
	local ti = {}
	for k,v in pairs(t) do
		ti[v] = k
	end
	return ti
end

function table.shuffle(t, from, to, random)
	from = from or 1
	to = to or #t
	random = random or math.random
	local n = to - from + 1
	while n > 1 do
		local r = from + n-1
		local l = from + random(0, n-1)
		t[l], t[r] = t[r], t[l]
		n = n-1
	end
end
