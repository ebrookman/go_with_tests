package arrays_and_slices

// Sum takes an array of numbers and returns the sum of them
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
