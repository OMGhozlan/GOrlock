package sliceoflife

import (
	"fmt"
)

// Slices returns a slice of length 5, with the elements set to 1 through 5 in order.
func Slices() (slice []int) {
	slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	return
}

func SumAllSlices(sliceA []int, sliceB []int) (sums []int) {
	sum := 0
	for _, number := range sliceA {
		sum += number
	}
	sums = append(sums, sum)
	sum = 0
	for _, number := range sliceB {
		sum += number
	}
	sums = append(sums, sum)
	return
}

// Sum takes an array of length 5 and returns the sum of its elements.
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

// SumAll takes a varying number of slices, and returns a slice of the sums of each of the passed slices.
func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}

// SumAllTails takes a varying number of slices, and returns a slice of the sums of the tail of each of the passed slices.
func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
