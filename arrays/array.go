package arrays

func Sum(numbers []int) int {
	sum := 0

	//range lets us itterate over an array, it returns index and its value
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// //make lets you create a slice of type first arg, with the length of 2nd arg
	// sums := make([]int, lengthOfNumbers)

	// //we loop over all
	// for i, number := range numbersToSum {
	// 	sums[i] = Sum(number)
	// }

	// return sums

	var sums []int

	for _, numbers := range numbersToSum {
		//append to the sums array with 2nd value
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
