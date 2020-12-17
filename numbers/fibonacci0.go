// Generate a Fibonacci sequence to the nth number

package main

import (
	"flag"
	"fmt"
)

func fibSequence(n int) {
	/* Generates a fibonacci sequence
	with the size of n */
	fib := []int{0, 1}
	for len(fib) <= n {
		fib = append(fib, fib[len(fib)-2]+fib[len(fib)-1])
	}
	fmt.Println(fib)
}

func main() {
	// Getting n as a flag
	var n int
	flag.IntVar(&n, "len", 7, "sequence length")
	flag.Parse()
	if n < 1 {
		fmt.Println("Error:", n, "is too low")
		return
	}
	fmt.Println(n)
	fibSequence(n)
}
