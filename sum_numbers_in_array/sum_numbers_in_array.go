package main

import "fmt"

func sum(arr []int) (result int64) {
	result = 0
	for i := 0; i < len(arr); i++ {
		result += int64(arr[i])
	}
	return
}

func main() {
	fmt.Println(sum([]int{1, 2, 3}))
}
