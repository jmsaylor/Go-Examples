package main

import "fmt"

func main()  {
	m := map[int]int{1:1}



	for i := 2; i <= 100; i++ {
		if m[i - 1] > 1000000 { 
			m[i] = 3 
		} else {
			m[i] = (m[i-1] * ((i % 6) + 1))
		}
	}

	// fmt.Println(m)


	for k, v := range(m) {
		fmt.Println(k * v)
	}

	variadic(5, 8, 15)

	doubler := closure()

	for i := 10; i >= 0; i-- {
		fmt.Println(doubler())
	}
}

func closure() func() int {
	var last int = 2
	return func() int {
		last *= 2
		return last
	}
}

func variadic(nums ...int)  {
	var product int = 1
	
	for _, num := range(nums) {
		fmt.Print(num, " ")
		product *= num
	}

	fmt.Println(" = ", product)

}