package main

import (
	"fmt"
)

func main() {
	var step int = 1
	fmt.Print("input banyak data: ")
	var i int
	fmt.Scanf("%d", &i)
	checkStep(i, &step)
	fmt.Println(step)
}

func checkStep(s int, step *int) {
	var x string
	for d := 0; d < s; d++ {
		fmt.Printf("input step U/D: ")
		fmt.Scanf("%s", &x)
		if x == "U" {
			*step++
		}
		if x == "D" {
			*step--
		}
	}
}
