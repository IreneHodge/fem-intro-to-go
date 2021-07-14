package main

import "fmt"

func average(nums...float64) float64 {
	total := 0.0
	for _, num := range nums {
       total += num
	}
	return total / float64(len(nums))
}

func main() {
	fmt.Println(average(10,5,7))
}