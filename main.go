package main

import "fmt"

// Average calculates the average of a list of floats
func Average(nums []float64) float64 {
	total := 0.0
	for _, x := range nums {
		total += x
	}
	return total / float64(len(nums))
}

func main() {
	numbers := []float64{1, 2, 3, 4}
	av := Average(numbers)
	fmt.Printf("Average is: %.2f", av)
}
