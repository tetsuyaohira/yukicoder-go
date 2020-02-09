package main

import (
	"fmt"
)

func main() {

	// 4 7 3
	// 3 2 1
	// 1 1000000000 500000000
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	cal1 := intAbs(b - a)
	cal2 := intAbs(b - c)
	cal3 := intAbs(a - c)

	fmt.Println(intMin(intMin(cal1, cal2), cal3))

}

func intAbs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func intMin(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
