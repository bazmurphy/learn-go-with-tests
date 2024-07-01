package arrays_and_slices

// Arrays have a fixed capacity which you define when you declare the variable.
// We can initialize an array in two ways:
// [N]type{value1, value2, ..., valueN} e.g. numbers := [5]int{1, 2, 3, 4, 5}
// [...]type{value1, value2, ..., valueN} e.g. numbers := [...]int{1, 2, 3, 4, 5}

// if you define a size it's an array
// Sum(numbers [5]int)

// if you don't define a size it's a slice
// Sum(numbers []int)

func Sum(numbers []int) int {
	sum := 0

	// default for loop
	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }

	// for loop with range operator
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// VARIADIC FUNCTIONS - functions that can take a variable number of arguments
// note the ... operator
func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)

	// this creates an empty slice, length 0, capacity 0
	// var sums []int

	// we can use `make` to create a slice of a specific length
	// it can also optionally take a capacityValue as a 3rd argument
	sums := make([]int, lengthOfNumbers)

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)

	// sums := make([]int, lengthOfNumbers)

	// for _, numbers := range numbersToSum {
	// 	tail := numbers[1:]
	// 	sums = append(sums, Sum(tail))
	// }

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
