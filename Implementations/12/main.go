/*
	Klavyeden sıfır girilene kadar alınan int türden sayılardan pozitif olanların ve negatif olanların ayrı ayrı kaç adet
	olduklarını ve toplamlarını bulan programı yazınız. Hiç pozitif ya da negatif sayı girilmediyse durumu ifade eden bilgilendirme
	çıktısını gerçekleştiriniz.
*/

package main

import "fmt"

func main() {
	//fmt.Printf("The total of your numbers is: %d", calculateTotal())
	var (
		nums                 []int
		negativeNums         []int
		positiveNums         []int
		numberOfNums         int
		numberOfNegativeNums int
		numberOfPositiveNums int
	)

	fmt.Print("- To terminate the process please insert \"0\" -\n\n")

	nums, numberOfNums = collectValues()
	negativeNums, numberOfNegativeNums = collectNegativeValues(nums)
	positiveNums, numberOfPositiveNums = collectPositiveValues(nums)

	if numberOfNums == 0 {
		fmt.Print("\n- ! There is no number ! -\n\n")
	}

	fmt.Printf("\n- Number of values: %d -\n", numberOfNums)
	printValues(nums)

	fmt.Printf("\n- Number of negative values: %d -\n", numberOfNegativeNums)
	printValues(negativeNums)

	fmt.Printf("\n- Number of positive values: %d -\n", numberOfPositiveNums)
	printValues(positiveNums)

	fmt.Printf("\n< Sum of numbers: %d >\n", sumOfValues(nums))
	fmt.Printf("\n< Sum of negative numbers: %d >\n", sumOfNegativeValues(negativeNums))
	fmt.Printf("\n< Sum of positive numbers: %d >\n", sumOfPositiveValues(positiveNums))
}

func collectValues() ([]int, int) {
	var value int = -1
	var arr []int = make([]int, 0)

	for value != 0 {
		fmt.Print("- Please insert an integer value: ")
		fmt.Scan(&value)

		if value != 0 {
			arr = append(arr, value)
		} else {
			break
		}
	}
	return arr, len(arr)
}

func collectNegativeValues(arr []int) ([]int, int) {
	negativeVals := make([]int, 0)

	for i := range arr {
		if arr[i] < 0 {
			negativeVals = append(negativeVals, arr[i])
		}
	}

	return negativeVals, len(negativeVals)
}

func collectPositiveValues(arr []int) ([]int, int) {
	positiveVals := make([]int, 0)

	for i := range arr {
		if arr[i] > 0 {
			positiveVals = append(positiveVals, arr[i])
		}
	}

	return positiveVals, len(positiveVals)
}

func printValues(arr []int) {
	if len(arr) != 0 {
		for i, val := range arr {
			fmt.Printf("* arr[%d] -> %d *\n", i, val)
		}
	}
}

func sumOfValues(arr []int) int {
	var total int = 0

	if len(arr) == 0 {
		return 0
	}

	for i := range arr {
		total = total + arr[i]
	}

	return total
}

func sumOfNegativeValues(arr []int) int {
	var total int = 0

	if len(arr) == 0 {
		return 0
	}

	for i := range arr {
		total = total + arr[i]
	}

	return total
}

func sumOfPositiveValues(arr []int) int {
	var total int = 0

	if len(arr) == 0 {
		return 0
	}

	for i := range arr {
		total = total + arr[i]
	}

	return total
}
