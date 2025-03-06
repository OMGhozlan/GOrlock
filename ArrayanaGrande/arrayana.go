package arrayanagrande

const arrayLength = 5

func ArraysOld() (array [arrayLength]int) {
	var arr [arrayLength]int
	for i := 0; i < arrayLength; i++ {
		arr[i] = i + 1
	}
	return
}

// Arrays returns an array of length 5, with the elements set to 1 through 5 in order.
func Arrays() (array [arrayLength]int) {
	for i := range array {
		array[i] = i + 1
	}
	return
}

// Sum takes an array of length 5 and returns the sum of its elements.
func Sum(numbers [arrayLength]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}
