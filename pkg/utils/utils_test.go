package utils_test

import (
	"testing"

	"github.com/AlexS778/myUtils/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestPopIntsStack(t *testing.T) {
	originalStack := []int{1, 2, 3, 4, 5}
	stack := []int{1, 2, 3, 4, 5}
	poppedValue := utils.PopFromIntsStack(&stack)
	assert.Equal(t, poppedValue, originalStack[len(originalStack)-1])
	assert.Equal(t, len(stack), len(originalStack)-1)
}

func TestMaxTwoInts(t *testing.T) {
	a, b := 10, 20
	maxInt := utils.MaxTwoInts(a, b)
	assert.Equal(t, b, maxInt)
	a, b = 20, 10
	maxInt = utils.MaxTwoInts(a, b)
	assert.Equal(t, a, maxInt)
}

func TestMinTwoInts(t *testing.T) {
	a, b := 10, 20
	minInt := utils.MinTwoInts(a, b)
	assert.Equal(t, a, minInt)
	a, b = 20, 10
	minInt = utils.MinTwoInts(a, b)
	assert.Equal(t, minInt, b)
}

func TestReverseInts(t *testing.T) {
	numbers := []int{2, 45, 3, 67, 87}
	reversedNumbers := utils.ReverseInts(numbers)
	assert.Equal(t, numbers, reversedNumbers)
}

func TestReverseString(t *testing.T) {
	got := "hello"
	want := "olleh"
	result := utils.ReverseString(got)
	assert.Equal(t, result, want)
}

func TestSortString(t *testing.T) {
	got := "hello"
	want := "ehllo"
	result := utils.SortString(got)
	assert.Equal(t, want, result)
}
