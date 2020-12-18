// Generate a Fibonacci sequence up to a number.
//If the number is not part of the sequence, the sequence is generated as long as the elements of the sequence are less than the given number.

package main

import (
	"flag"
	"fmt"
)

func fibSequence(n int) {
	// Generates a fibonacci sequence up to n
	i, j := 0, 1
	for true {
		if j > n {
			fmt.Println(i)
			return
		i, j = j, i+j
	}
}

func main() {
	// Getting n as a flag
	var n int
	flag.IntVar(&n, "num", 7, "largest number")
	flag.Parse()
	if n < 0 {
		fmt.Println("Error:", n, "is too low")
		return
	} else if n == 0 || n == 1 {
		fmt.Println(n)
		return
	}
	fibSequence(n)
}
