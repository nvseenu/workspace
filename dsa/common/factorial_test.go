package common

import (
	"testing"
)

var factorialTestData = []struct {
	input  uint64
	expRes uint64
}{
	{input: 0, expRes: 1},
	{input: 1, expRes: 1},
	{input: 2, expRes: 2},
	{input: 5, expRes: 120},
	{input: 13, expRes: 6227020800},
}

func TestFactorial(t *testing.T) {
	for _, td := range factorialTestData {
		res := Factorial(td.input)
		if td.expRes != res {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
