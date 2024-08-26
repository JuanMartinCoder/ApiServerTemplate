package utils

type Routes struct {
	MAIN  string
	ERROR string
}

var RoutesIntance = Routes{
	MAIN:  "/",
	ERROR: "/error",
}
