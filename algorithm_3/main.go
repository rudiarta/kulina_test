package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//1345679
	fmt.Print("input nilai: ")
	var i int
	fmt.Scanf("%d", &i)
	extractInt(i)
}

func extractInt(n int) {
	slice := strconv.Itoa(n)
	var i int = len(slice) - 1
	for _, v := range slice {
		compare := math.Pow(10, float64(i))
		value := int(v-'0') * int(compare)
		fmt.Println(value)
		i--
	}
}
