// Generate a Fibonacci sequence to the nth number

package main

import (
	"fmt"
)

func fibSequence(n int) {
	/*Generates a fibonacci sequence
	with the size of n*/
	fib := []int{0, 1}
	for len(fib) <= n {
		fib = append(fib, fib[len(fib)-2]+fib[len(fib)-1])
	}
	fmt.Println(fib)
}

func main() {
	var n int
	fmt.Println("How many numbers do you need?")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n < 1 {
		fmt.Println("Error:", n, "is too low")
		return
	}
	fibSequence(n)
}
