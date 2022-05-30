# Hydra API documentation

## Lua version
Hydra uses gopher-lua, Lua 5.1

## Globals

- `arg`: table containing command line arguments
- `hydra`: contains minetest protocol API, see [hydra.md](hydra.md)
- `vec2`, `vec3`, `box`: vector library, see [vector.md](vector.md)

## Additional packages

`require()` can be used to import these modules.

- `escapes`: contains utility functions to deal with minetest escape sequences, see [escapes.md](escapes.md)
- `client`: a function to create a client from command line arguments in the form `<server> <username> <password>`. This is trivial but so commonly used that this function was added to avoid repetition in scripts.

## Standard library additions

Source code: [luax](https://github.com/EliasFleckenstein03/luax).

- `table.indexof(list, val)`
- `table.copy(t, seen)`
- `table.insert_all(t, other)`
- `table.key_value_swap(t)`
- `table.shuffle(t, from, to, random)`
- `string.split(str, delim, include_empty, max_splits, sep_is_pattern)`
- `string.trim(str)`
- `math.hypot(x, y)`
- `math.sign(x, tolerance)`
- `math.factorial(x)`
- `math.round(x)`
- `math.clamp(min, max, v)`
