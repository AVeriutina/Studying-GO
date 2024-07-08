package main

import "fmt"

func revese_str(str string) string {
	var rev_str = ""
	for i := len(str) - 1; i >= 0; i-- {
		rev_str += string(str[i])
	}
	return rev_str
}

func main() {
	fmt.Println(revese_str("hello world"))
}
