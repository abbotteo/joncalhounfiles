package testing101

func Sum(numbers []int) int {
	sum := 0
	// This bug is intentional
	for _, n := range numbers {
		sum += n
	}
	return sum
}
