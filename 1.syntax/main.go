package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
	j := 10
	for ; j > 0; j-- {
		fmt.Println(j)
	}
}
