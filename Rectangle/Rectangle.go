package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func calc_sq(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	r := Rectangle{width: 3, height: 4}
	fmt.Println("area: ", calc_sq(r))
}
