package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	fmt.Scan(&n)

	// 1 → 1
	// 2 → 1122
	// 3 → 111222333
	// 4 → 1111222233334444
	// 5 → 1111122222333334444455555

	var result string
	for i := 1; i <= n; i++ {
		for j := 0; j < n; j++ {

			var str = strconv.Itoa(i)
			result += str[len(str)-1:]
		}
	}

	fmt.Println(result)
}
