package main

import (
	"clay-ui/clay"
	"fmt"
)

func main() {
	fmt.Println("size", clay.MinMemorySize())

	a := clay.NewArena()
	handle := func(err error) {
		panic(err)
	}

	clay.Initialize(a, clay.Dimensions{Width: 200, Height: 200}, handle)
	// causes segfault for some reason
	// clay.BeginLayout()
}
