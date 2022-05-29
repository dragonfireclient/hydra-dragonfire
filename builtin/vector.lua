--[[ builtin/vector.lua ]]--

local function wrap(op, body_wrapper, ...)
	return loadstring("return function(a, b) " .. body_wrapper(op, ...) .. "end")()
end

local function arith_mt(...)
	return {
		__add = wrap("+", ...),
		__sub = wrap("-", ...),
		__mul = wrap("*", ...),
		__div = wrap("/", ...),
		__mod = wrap("%", ...),
	}
end

-- vec2

local mt_vec2 = arith_mt(function(op)
	return [[
		if type(b) == "number" then
			return vec2(a.x ]] .. op.. [[ b, a.y ]] .. op .. [[ b)
		else
			return vec2(a.x ]] .. op.. [[ b.x, a.y ]] .. op.. [[ b.y)
		end
	]]
end)

function mt_vec2:__neg()
	return vec2(-self.x, -self.y)
end

function mt_vec2:__tostring()
	return "(" .. self.x .. ", " .. self.y .. ")"
end

mt_vec2.__index = {
	validate = function(self)
		assert(type(self.x) == "number")
		assert(type(self.y) == "number")
		return self
	end
}

function vec2(a, b)
	local o = {}

	if type(a) == "number" then
		o.x = a
		o.y = b or a
	else
		o.x = a.x
		o.y = a.y
	end

	setmetatable(o, mt_vec2)
	return o:validate()
end

-- vec3

local mt_vec3 = arith_mt(function(op)
	return [[
		if type(b) == "number" then
			return vec3(a.x ]] .. op.. [[ b, a.y ]] .. op .. [[ b, a.z ]] .. op .. [[ b)
		else
			return vec3(a.x ]] .. op.. [[ b.x, a.y ]] .. op.. [[ b.y, a.z ]] .. op.. [[ b.z)
		end
	]]
end)

function mt_vec3:__neg()
	return vec3(-self.x, -self.y, -self.z)
end

function mt_vec3:__tostring()
	return "(" .. self.x .. ", " .. self.y .. ", " .. self.z .. ")"
end

mt_vec3.__index = {
	validate = function(self)
		assert(type(self.x) == "number")
		assert(type(self.y) == "number")
		assert(type(self.z) == "number")
		return self
	end
}

function vec3(a, b, c)
	local o = {}

	if type(a) == "number" then
		o.x = a
		o.y = b or a
		o.z = c or a
	else
		o.x = a.x
		o.y = a.y
		o.z = a.z
	end

	setmetatable(o, mt_vec3)
	return o:validate()
end

-- box

local mt_box = arith_mt(function(op)
	return "return box(a.min " .. op .. " b, a.max " .. op .. " b)"
end)

function mt_box:__neg()
	return box(-self.min, -self.max)
end

function mt_box:__tostring()
	return "[" .. self.min .. "; " .. self.max .. "]"
end

mt_box.__index = {
	contains = function(a, b)
		if type(b) == "number" or b.x then
			return a.min <= b and a.max >= b
		else
			return a.min <= b.min and a.max >= b.max
		end
	end,
	validate = function(self)
		if type(self.min) == "number" then
			assert(type(self.max) == "number")
		else
			assert(not self.min.z == not self.max.z)
			self.min:validate()
			self.max:validate()
		end
	end,
}

function box(a, b)
	local o = {}

	if type(a) == "number" or a.x then
		o.min = a
		o.max = b
	else
		o.min = a.min
		o.max = a.max
	end

	setmetatable(o, mt_box)
	return o:validate()
end
