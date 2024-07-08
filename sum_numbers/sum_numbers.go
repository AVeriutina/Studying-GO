package main

import "fmt"

func sum(a int64, b int64) int64 {
	return a + b
}

func main() {
	var x, y int64
	fmt.Scanf("%d %d", &x, &y)

	fmt.Println(sum(x, y))
}
