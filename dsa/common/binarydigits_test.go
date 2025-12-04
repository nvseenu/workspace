package common

import (
	"fmt"
	"slices"
	"testing"
)

var binaryDigitsTestData = []struct {
	input  int
	expRes [][]byte
}{
	{input: 1, expRes: [][]byte{
		{0},
		{1},
	}},
	{input: 2, expRes: [][]byte{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}},
	{input: 3, expRes: [][]byte{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1},
	}},
}

func TestBinaryDigits(t *testing.T) {
	for _, td := range binaryDigitsTestData {
		fmt.Println("TestData:", td)
		res := BinaryDigits(td.input)
		for i := 0; i < len(td.expRes); i++ {
			if !slices.Equal(td.expRes[i], res[i]) {
				t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes[i], res[i], td)
			}
		}
	}
}
