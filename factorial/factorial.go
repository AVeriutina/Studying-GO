package main

import "fmt"

func factorial(n uint64) (result uint64) {
	result = 1
	var i uint64
	for i = 2; i <= n; i++ {
		result *= i
	}
	return
}

func main() {
	var n uint64
	fmt.Scanf("%d", &n)

	fmt.Println(factorial(n))
}
