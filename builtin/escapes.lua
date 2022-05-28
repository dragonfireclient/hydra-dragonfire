--[[ builtin/escapes.lua ]]--
-- code taken from minetest/builtin/common/misc_helpers.lua with modifications

local escapes = {}
package.loaded["escapes"] = escapes

escapes.ESCAPE_CHAR = string.char(0x1b)

function escapes.get_color_escape_sequence(color)
	return escapes.ESCAPE_CHAR .. "(c@" .. color .. ")"
end

function escapes.get_background_escape_sequence(color)
	return escapes.ESCAPE_CHAR .. "(b@" .. color .. ")"
end

function escapes.colorize(color, message)
	local lines = tostring(message):split("\n", true)
	local color_code = escapes.get_color_escape_sequence(color)

	for i, line in ipairs(lines) do
		lines[i] = color_code .. line
	end

	return table.concat(lines, "\n") .. escapes.get_color_escape_sequence("#ffffff")
end

function escapes.strip_foreground_colors(str)
	return (str:gsub(escapes.ESCAPE_CHAR .. "%(c@[^)]+%)", ""))
end

function escapes.strip_background_colors(str)
	return (str:gsub(escapes.ESCAPE_CHAR .. "%(b@[^)]+%)", ""))
end

function escapes.strip_colors(str)
	return (str:gsub(escapes.ESCAPE_CHAR .. "%([bc]@[^)]+%)", ""))
end

function escapes.strip_translations(str)
	return (str
		:gsub(escapes.ESCAPE_CHAR .. "%(T@[^)]+%)", "")
		:gsub(escapes.ESCAPE_CHAR .. "[TFE]", ""))
end

function escapes.strip_all(str)
	str = escapes.strip_colors(str)
	str = escapes.strip_translations(str)
	return str
end
