package main

import (
	"fmt"
)

func main() {
	fmt.Print("input banyak data: ")
	var i int
	fmt.Scanf("%d", &i)
	data := sockMerchant(i)
	fmt.Println(data)
}

func sockMerchant(n int) int {
	var data []int
	var i int = 0
	var v int
	for i < n {
		v = 0
		fmt.Print("input nilai: ")
		fmt.Scanf("%d", &v)
		data = append(data, v)
		i++
	}

	return checkData(data)
}

func checkData(d []int) int {
	var i int = 0
	var tmp []int = d
	var result int = 0
	for i < len(tmp) {
	outerLoop:
		for index, value := range tmp {
			for x, y := range tmp[index+1:] {
				if value == y {
					tmp = append(tmp[:index], tmp[index+1:]...)
					tmp = append(tmp[:x], tmp[x+1:]...)
					result++
					break outerLoop
				}
			}
		}
		i++
	}

	return result
}
