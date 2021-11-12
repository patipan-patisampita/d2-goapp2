package main

import "fmt"

var (
	contry = "Thailand" //global variable
	province = "trat"
	poscode = 23000
)

func main() {
	y := false //local varible
	fmt.Println(y)

	fmt.Println(contry)
	fmt.Println(province)
	fmt.Println(poscode)
}
