// Find prime numbers until the user chooses to stop asking for the next one.

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func isPrime(n float64) bool {
	// The simplest check to see if a number belongs to prime numbers.
	for i := 2.0; i <= math.Sqrt(n); i++ {
		if int(n)%int(i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input := "\n"
	for i := 2.0; input == "\n"; i++ {
		if isPrime(i) == true {
			fmt.Println(i)
			fmt.Println("Press enter for next prime - enter stop to end")
			input, _ = reader.ReadString('\n')
		}
	}
}
