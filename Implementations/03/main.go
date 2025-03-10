package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func main() {
	var p1, p2 Point

	fmt.Print("Enter coordinates for 1st point [x1, y1]: ")
	_, err := fmt.Scan(&p1.x, &p1.y)
	if err != nil {
		fmt.Println("Invalid input for the first point. Please enter two float numbers.")
		return
	}

	fmt.Print("Enter coordinates for 2nd point [x2 y2]: ")
	_, err = fmt.Scan(&p2.x, &p2.y)
	if err != nil {
		fmt.Println("Invalid input for the second point. Please enter two float numbers.")
		return
	}

	distance := euclideanDistance(p1, p2)
	fmt.Printf("Euclidean Distance between {%.2f, %.2f} and {%.2f, %.2f} is: %.4f\n", p1.x, p1.y, p2.x, p2.y, distance)
}

func euclideanDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}
