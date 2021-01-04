// Write a program that prints the numbers 1 to n.
// For multiples of three, print "Fizz".
// For multiples of five, print "Buzz".
// For multiples of both three and five, print "FizzBuzz".

package main

import (
	"flag"
	"fmt"
)

func fizzbuzz(num int) string {
	if num%3 == 0 {
		if num%5 == 0 {
			return "FizzBuzz"
		} else {
			return "Fizz"
		}
	} else if num%5 == 0 {
		return "Buzz"
	} else {
		return "num"
	}
}

func main() {
	var n int
	flag.IntVar(&n, "num", 100, "largest number")
	flag.Parse()
	if n <= 0 {
		fmt.Println("Error:", n, "is too low")
		return
	}
	for i := 1; i <= n; i++ {
		j := fizzbuzz(i)
		if j == "num" {
			fmt.Println(i)
		} else {
			fmt.Println(j)
		}
	}
}
