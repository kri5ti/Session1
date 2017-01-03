package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	for i, p := range primes {
		i++
		fmt.Println("The", i, "st prime is", p)
	}
}
