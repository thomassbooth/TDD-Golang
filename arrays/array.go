package arrays

func Sum(numbers []int) int {
	sum := 0

	//range lets us itterate over an array, it returns index and its value
	for _, number := range numbers {
		sum += number
	}

	return sum
}
