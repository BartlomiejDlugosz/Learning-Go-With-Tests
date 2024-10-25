package arrays_and_slices

func Sum(numbers [5]int) interface{} {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
