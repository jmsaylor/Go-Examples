package main

import (
	"fmt"
)

type brick struct {
	x int
	y int
}

func makeBrick(x int, y int) *brick {
	b := brick{x : x, y : y}
	return &b
}

func main()  {
	
	bricks := make([]*brick, 0)

	for i := 0; i < 100; i++ {
		bricks = append(bricks, makeBrick(i / 10, i % 10))
	}

	for i := 0; i < 100; i++ {
		fmt.Println(bricks[i])
	}
	
}