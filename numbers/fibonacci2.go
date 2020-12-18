// Get the nth number of the Fibonacci sequence

package main

import (
	"flag"
	"fmt"
)

func fibSequence(n int) {
	// Generates a fibonacci sequence up to n
	i, j := 0, 1
	for k := 2; k < n; k++ {
		i, j = j, i+j
	}
	fmt.Println(j)
}

func main() {
	// Getting n as a flag
	var n int
	flag.IntVar(&n, "l", 7, "sequence length")
	flag.Parse()
	if n < 0 {
		fmt.Println("Error:", n, "is too low")
		return
	}

	fibSequence(n)
}
