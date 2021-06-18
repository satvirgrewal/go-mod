package calc

func Add(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}
