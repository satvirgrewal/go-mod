package calc

import "errors"

func Add(nums ...int) (error, int) {
	sum := 0
	if len(nums) < 2 {
		err := errors.New("")
		return err, sum
	}
	for _, val := range nums {
		sum += val
	}
	return nil, sum
}
