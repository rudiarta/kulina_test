package main

import "fmt"

func main() {
	var step int = 1
	var lamp [100]bool
	for i := 0; i < 100; i++ {
		stepAction(step, &lamp)
		step++
		if step > 3 {
			step = 1
		}
	}

	fmt.Println(checkLamp(&lamp))
}

func stepAction(step int, l *[100]bool) {
	if step == 1 {
		for i, _ := range l {
			l[i] = true
		}
	} else if step == 2 {
		for i, _ := range l {
			if ((i + 1) % 2) == 0 {
				l[i] = true
			} else {
				l[i] = false
			}
		}
	} else if step == 3 {
		for i, _ := range l {
			if ((i + 1) % 3) == 0 {
				l[i] = true
			} else {
				l[i] = false
			}
		}
	}
}

func checkLamp(l *[100]bool) int {
	var tmp = 0
	for _, v := range l {
		if v == true {
			tmp++
		}
	}

	return tmp
}
