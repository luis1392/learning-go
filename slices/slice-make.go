package main

import (
	"fmt"
)

func main()  {
	x := make([]int, 10, 100)
	x[0] = 1
	x[1] = 10
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 200)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}