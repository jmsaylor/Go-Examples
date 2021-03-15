package main

import (
	"fmt"
	"time"
)

func main()  {

	var words = []string{"Mars", "Jupiter"}
	
	for i := 0; i < 10; i++ {
		x, y := test(i, words)
		fmt.Println(y, x)
		var date = time.Now()
		fmt.Print(date)
		fmt.Print(" ")
		if i % 2 == 0 {
			fmt.Println(i)
		} else {
			fmt.Print(i)
			fmt.Println(" beep")
		}
	}
}

func test(a int, b []string) (string, int){
	return b[a%2], a
}