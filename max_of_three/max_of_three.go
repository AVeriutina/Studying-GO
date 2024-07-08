package main

import "fmt"

func maxOfThree(a, b, c int64) (result int64) {
	result = a
	if b > result {
		result = b
	}
	if c > result {
		result = c
	}
	return
}

func main() {
	var x, y, z int64
	fmt.Scanf("%d %d %d", &x, &y, &z)

	fmt.Println(maxOfThree(x, y, z))
}
