-- hello.lua
-- the first program in every language

io.write("Hello world, from ",_VERSION,"!\n")

local m = require("mymodule")
m.myfunc()
print(m.name)