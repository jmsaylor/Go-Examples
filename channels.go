package main

import (
	"fmt"
)

func main()  {
	c := make(chan int)
	q := make(chan int, 1) 

	go func ()  {
		for i := 0; i < 50; i++ {
			c <- i
		}
		q <- 7
	}()
	fibonacci(c, q)
}

func fibonacci(c, q chan int)  {
	x, y := 0, 1
	for {
		select {
		case <- c:
			x, y = y, (x + y)
			fmt.Print(y, " ")
		case <- q:
			fmt.Println("done")
			return
		}
	
	}
}