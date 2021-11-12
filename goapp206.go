package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, dept = "Trattc", "IT"
	s := fmt.Sprintln(name,dept)
	io.WriteString(os.Stdout,s)
}
