package main

import (
	"fmt"
)

func main()  {
	//map[string]int usar map: [string] la llave o identificador de tipo string y int el valor 
	m := map[string]int{
		"Batman": 40,
		"Robin": 17
	}
	fmt.Println(m)

	fmt.Println(m["Batman"])

	// valor , ok -> en leguaje idiomatico de go si un valor exste o no 
	v, ok := m["Luis"]
	fmt.Println(v)
	fmt.Println(ok)

	//Validar si un tiene una llave o identificador  
	if v, ok := m["Luis"]; ok{
		fmt.Println("Imprimiendo desde el if", v)
	}

}