module firefox123456/mypackege2

go 1.19

require (
   "firefox123456/mypackege" v0.0.0
)

replace (
	"firefox123456/mypackege" => "./../mypackege"
)