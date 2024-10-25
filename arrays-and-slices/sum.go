package arrays_and_slices

func Sum(numbers []int) interface{} {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
