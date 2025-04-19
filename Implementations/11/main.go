package main

import "fmt"

func main() {
	fmt.Printf("The total of your numbers is: %d", calculateTotal())
}

func calculateTotal() int {
	var value int
	var check string

	total := 0
	check = "y"

	for true && check == "y" {
		fmt.Print("Insert a number: ")
		fmt.Scan(&value)

		if value == 0 {
			return total
		}

		total += value

		fmt.Print("Do you want to continue? [y/n]: ")
		fmt.Scan(&check)
	}
	return total
}
