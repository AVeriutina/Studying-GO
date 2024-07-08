package main

import "fmt"

func isEven(a int64) bool {
	return a%2 == 0
}

func main() {
	var x int64
	fmt.Scanf("%d", &x)

	if isEven(x) {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
