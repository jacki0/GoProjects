// Generate a Fibonacci sequence up to a number.
//If the number is not part of the sequence, the sequence is generated as long as the elements of the sequence are less than the given number.

package main

import (
	"flag"
	"fmt"
)

func fibSequence(n int) {
	// Generates a fibonacci sequence up to n
	fib := []int{0, 1}
	var i int
	for true {
		i = fib[len(fib)-2] + fib[len(fib)-1]
		if i > n {
			break
		}
		fib = append(fib, i)
	}
	fmt.Println(fib)
}

func main() {
	// Getting n as a flag
	var n int
	flag.IntVar(&n, "num", 7, "largest number")
	flag.Parse()
	if n < 0 {
		fmt.Println("Error:", n, "is too low")
		return
	}
	fibSequence(n)
}
