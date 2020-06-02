package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	z := []int{9, 99, 999}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println("")

	// dividir slice
	fmt.Println(x[1:3])
	fmt.Println(x[2:4])

	for i, v := range x {
		fmt.Println(i, " ", v)
	}

	// agregar valores a un slice
	y := append(x, 44, 55, 66 )
	fmt.Println(y)

	// concatenar otro slice
	y = append(y, z...)
	fmt.Println(y)

	// quitando elementos de un slice
	y = append(y[:2], x[4:]...)
	fmt.Println(y)

}
