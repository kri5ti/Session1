package main

import (
	"fmt"
	"regexp"
)

func main() {
	// These are the primes less than 200
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29,
		31, 37, 41, 43, 47, 53, 59, 61, 67, 71,
		73, 79, 83, 89, 97, 101, 103, 107, 109,
		113, 127, 131, 137, 139, 149, 151, 157,
		163, 167, 173, 179, 181, 191, 193, 197, 199}
	fmt.Println(primes)
	// Write a program to print only the primes less than 10
	var a []int

	for _, v := range primes {
		if v < 10 {
			a = append(a, v)
		}
	}
	fmt.Println("Slice with less than 10 values:\n", a)
	// Bonus: write a print only the two digit primes
	reg := regexp.MustCompile(`^[0-9][0-9]$`)
	fmt.Println("two digit primes: ")
	for _, v := range primes {
		if reg.MatchString(fmt.Sprintf("%d", v)) {
			fmt.Print(v, " ")
		}
	}

	slice := primes[4:25]
	fmt.Println("\nSlice with only two digit primes:\n", slice)

}
