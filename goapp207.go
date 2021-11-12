package main

import "fmt"

func main() {
	var a float64 = 90.0 //Static type
	a = 50.5
	var b string = "welcome";
	y := 14  //Dynamic Type 

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(y)
	fmt.Printf("A is of type %T\n", a)
	fmt.Printf("A is of type %T", y)
}
