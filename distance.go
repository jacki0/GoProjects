// Calculate the distance between two points that are given by coordinates

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func calculate(coord []float64) float64 {
	const earth_rad float64 = 6372795

	// converting coordinates to radians
	lat1 := coord[0] * math.Pi / 180
	long1 := coord[1] * math.Pi / 180
	lat2 := coord[2] * math.Pi / 180
	long2 := coord[3] * math.Pi / 180

	// Obtaining cosines and sines of latitudes and differences of longitudes
	cos_lat1 := math.Cos(lat1)
	sin_lat1 := math.Sin(lat1)
	cos_lat2 := math.Cos(lat2)
	sin_lat2 := math.Sin(lat2)
	dellta := long2 - long1
	cos_delta := math.Cos(dellta)
	sin_delta := math.Sin(dellta)

	// Calculating the length of a large circle
	y := math.Sqrt(math.Pow(cos_lat2*sin_delta, 2) + math.Pow(cos_lat1*sin_lat2-sin_lat1*cos_lat2*cos_delta, 2))
	x := sin_lat1*sin_lat2 + cos_lat1*cos_lat2*cos_delta

	// Calculating the angular difference
	ang_dif := math.Atan2(y, x)
	distance := ang_dif * earth_rad
	return distance
}

func main() {
	// Getting coordinates
	var coord []float64
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
	input := strings.Split(point1+" "+point2, " ")
	for _, s := range input {
		num, err := strconv.ParseFloat(s, 64)
		if err == nil {
			coord = append(coord, num)
		}
	}
	calculate(coord)
}
