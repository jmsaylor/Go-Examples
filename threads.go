package main

import (
	"fmt"
	"time"
)

func main()  {
	counter1 := counter()
	counter2 := counter()

	go counter1()
	go counter2()

	time.Sleep(time.Second)

}

func counter() func() {
	return func() {
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
	}

}