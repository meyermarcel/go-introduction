package sum

import "fmt"

type Test struct {
}

type ErrSum struct {
	message string
}

func (e *ErrSum) Error() string {
	return e.message
}

// This sums two numbers
func SumPositive(value1, value2 int) (error, int) {
	if value1 < 0 || value2 < 0 {

		return &ErrSum{message: "negative number"}, 0
	}
	number := 0

	return nil, value1 + value2

}
