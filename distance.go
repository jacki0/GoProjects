package main

import (
	"fmt"
)

func main() {
	// Getting coordinates
	var point1, point2 string
	fmt.Println("Enter the coordinates of the first point")
	_, err := fmt.Scanf("%s", &point1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Enter the coordinates of the second point")
	_, err = fmt.Scanf("%s", &point2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
