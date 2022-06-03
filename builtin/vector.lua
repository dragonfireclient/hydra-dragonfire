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
		__index = {},
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

function mt_vec2:__len()
	return math.sqrt(self:dot(self))
end

function mt_vec2:__equ(other)
	return type(other) == "table" and self.x == other.x and self.y == other.y
end

function mt_vec2:__tostring()
	return "(" .. self.x .. ", " .. self.y .. ")"
end

function mt_vec2.__index:validate()
	assert(type(self.x) == "number")
	assert(type(self.y) == "number")
	return self
end

function mt_vec2.__index:round()
	return vec2(math.round(self.x), math.round(self.y))
end

function mt_vec2.__index:manhattan()
	return math.abs(self.x) + math.abs(self.y)
end

function mt_vec2.__index:volume()
	return self.x * self.y
end

function mt_vec2.__index:dot(other)
	return self.x * other.x + self.y * other.y
end

function mt_vec2.__index:norm()
	return self / #self
end

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

function mt_vec3:__len()
	return math.sqrt(self:dot(self))
end

function mt_vec3:__equ(other)
	return type(other) == "table" and self.x == other.x and self.y == other.y and self.z == other.z
end

function mt_vec3:__tostring()
	return "(" .. self.x .. ", " .. self.y .. ", " .. self.z .. ")"
end

function mt_vec3.__index:validate()
	assert(type(self.x) == "number")
	assert(type(self.y) == "number")
	assert(type(self.z) == "number")
	return self
end

function mt_vec3.__index:round()
	return vec3(math.floor(self.x), math.floor(self.y), math.floor(self.z))
end

function mt_vec3.__index:manhattan()
	return math.abs(self.x) + math.abs(self.y) + math.abs(self.z)
end

function mt_vec3.__index:volume()
	return self.x * self.y * self.z
end

function mt_vec3.__index:dot(other)
	return self.x * other.x + self.y * other.y + self.z * other.z
end

function mt_vec3.__index:cross(other)
	return vec3(
		self.y * other.z - self.z * other.y,
		self.z * other.x - self.x * other.z,
		self.x * other.y - self.y * other.x
	)
end

function mt_vec3.__index:norm()
	return self / #self
end

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

function mt_box:__eq(other)
	return self.min == other.min and self.max == other.max
end

function mt_box:__tostring()
	return "[" .. self.min .. "; " .. self.max .. "]"
end

function mt_box.__index:validate()
	if type(self.min) == "number" then
		assert(type(self.max) == "number")
	else
		assert(not self.min.z == not self.max.z)
		self.min:validate()
		self.max:validate()
	end
end

function mt_box.__index:volume()
	local diff = self.max - self.min
	if type(diff) == "number" then
		return diff
	else
		return diff:volume()
	end
end

function mt_box.__index:contains(other)
	if type(other) == "number" or other.x then
		return self.min <= other and self.max >= other
	else
		return self.min <= other.min and self.max >= other.max
	end
end

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
