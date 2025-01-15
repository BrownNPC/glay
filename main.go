package main

import (
	"clay-ui/clay"
	"fmt"
)

func main() {
	fmt.Println(clay.MinMemorySize())

	a := clay.NewArena()
	handle := func(err error) {
		fmt.Println(err)
	}

	clay.Initialize(a, clay.Dimensions{Width: 200, Height: 200}, handle)

}
