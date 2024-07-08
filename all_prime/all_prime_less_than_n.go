package main

import "fmt"

func all_prime(n int) (primes []int) {
	lis := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if lis[i] {
			continue
		}
		for j := 2 * i; j <= n; j += i {
			lis[j] = true
		}
	}
	for i := 2; i < n+1; i++ {
		if !lis[i] {
			primes = append(primes, i)
		}
	}
	return
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(all_prime(n))

}
