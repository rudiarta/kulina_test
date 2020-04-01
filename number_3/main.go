package main

import (
	"fmt"
)

func main() {
	var input int = 1345679
	data := input % 10000
	data = (input - data) - 1000000
	fmt.Println(data)
}
