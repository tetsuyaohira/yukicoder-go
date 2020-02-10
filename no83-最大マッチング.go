package main

import "fmt"

// 2 1
// 3 7
// 4 11
// 5 17
// 6 111
// 7 117
// 8 1111
// 9 1117
// 10 11111

// 20 1111111111
func main() {
	var n int
	fmt.Scan(&n)
	// 偶数の場合、2で割った数字ぶん1
	// 奇数の場合、2で割った数字ぶん1+7
	var result string
	if n == 3 {
		result = "7"
	} else if n%2 == 0 {
		// 偶数
		for i := 0; i < n/2; i++ {
			result += "1"
		}
	} else {
		// 奇数
		result += "7"
		for i := 0; i < (n-3)/2; i++ {
			result += "1"
		}
	}
	fmt.Println(result)
}
