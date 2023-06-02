package utils

import "sort"

// The PopFromIntsStack function is used to retrieve and remove the top element from a stack implemented as a slice of integers.
// It takes a pointer to the stack as its parameter and returns the value of the top element that is being removed.
//
// Parameters:
//   - stack *[]int: A pointer to a stack implemented as a slice of integers. The stack should be passed as a pointer to enable modifications to the original stack.
//
// Return Value:
//   - int: The value of the top element in the stack that is being removed.
//
// Side Effects:
//   - The PopFromIntsStack function modifies the original stack by removing the top element.
//
// Usage:
//
//	To use the PopFromIntsStack function, you should provide a pointer to the stack as an argument.
//	The function will return the value of the top element and remove it from the stack.
//
//	stack := []int{1, 2, 3, 4, 5}
//	topElement := PopFromIntsStack(&stack)
//	fmt.Println(topElement) // Output: 5
//	fmt.Println(stack)      // Output: [1 2 3 4]
//
// Note:
//
//	It is important to ensure that the stack is not empty before calling the PopFromIntsStack function,
//	as it does not perform any checks for an empty stack. Attempting to pop from an empty stack will result in a runtime error.
func PopFromIntsStack(stack *[]int) int {
	defer func() {
		*stack = (*stack)[:len(*stack)-1]
	}()
	return (*stack)[len(*stack)-1]
}

// The MaxTwoInts function is used to determine the maximum value between two integers.
// It takes two integers, 'a' and 'b', as parameters and returns the larger value.
//
// Parameters:
//   - a int: The first integer to compare.
//   - b int: The second integer to compare.
//
// Return Value:
//   - int: The maximum value between 'a' and 'b'.
//
// Usage:
//
//	To use the MaxTwoInts function, provide two integers as arguments.
//	The function will return the larger value.
//
//	result := MaxTwoInts(3, 7)
//	fmt.Println(result) // Output: 7
//
// Note:
//
//	This function assumes that the inputs are integers.
//	If 'a' and 'b' have equal values, the function will return either 'a' or 'b'.
func MaxTwoInts(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// The MinTwoInts function is used to determine the minimum value between two integers.
// It takes two integers, 'a' and 'b', as parameters and returns the smaller value.
//
// Parameters:
//   - a int: The first integer to compare.
//   - b int: The second integer to compare.
//
// Return Value:
//   - int: The minimum value between 'a' and 'b'.
//
// Usage:
//
//	To use the MinTwoInts function, provide two integers as arguments.
//	The function will return the smaller value.
//
//	result := MinTwoInts(3, 7)
//	fmt.Println(result) // Output: 3
//
// Note:
//
//	This function assumes that the inputs are integers.
//	If 'a' and 'b' have equal values, the function will return either 'a' or 'b'.
func MinTwoInts(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// The ReverseInts function is used to reverse a slice of integers in-place.
// It takes a slice of integers, 'nums', as a parameter and modifies the slice to reverse the order of its elements.
// The function returns the reversed slice.
//
// Parameters:
//   - nums []int: A slice of integers to be reversed.
//
// Return Value:
//   - []int: The reversed slice of integers.
//
// Usage:
//
//	To use the ReverseInts function, provide a slice of integers as an argument.
//	The function will modify the original slice to reverse the order of its elements.
//	It will also return the reversed slice for convenience.
//
//	numbers := []int{1, 2, 3, 4, 5}
//	reversed := ReverseInts(numbers)
//	fmt.Println(reversed) // Output: [5 4 3 2 1]
//	fmt.Println(numbers)  // Output: [5 4 3 2 1]
//
// Note:
//
//	This function modifies the original slice in-place, so it does not create a new slice.
//	The function utilizes an efficient swapping technique to reverse the elements.
func ReverseInts(nums []int) []int {
	length := len(nums)
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
	return nums
}

// The ReverseString function is used to reverse a string.
// It takes a string, 's', as a parameter and returns the reversed string.
//
// Parameters:
//   - s string: The string to be reversed.
//
// Return Value:
//   - string: The reversed string.
//
// Usage:
//
//	To use the ReverseString function, provide a string as an argument.
//	The function will return the reversed version of the input string.
//
//	str := "Hello, world!"
//	reversed := ReverseString(str)
//	fmt.Println(reversed) // Output: "!dlrow ,olleH"
//
// Note:
//
//	This function creates a new string by iterating through the characters of the input string in reverse order.
func ReverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}
	return string(runes)
}

// The SortString function is used to sort the characters of a string in ascending order.
// It takes a string, 's', as a parameter and returns a new string with the characters sorted.
//
// Parameters:
//   - s string: The string to be sorted.
//
// Return Value:
//   - string: The sorted string.
//
// Usage:
//
//	To use the SortString function, provide a string as an argument.
//	The function will return a new string with the characters sorted in ascending order.
//
//	str := "Hello"
//	sorted := SortString(str)
//	fmt.Println(sorted) // Output: "olleH"
//
// Note:
//
//	This function creates a new string by converting the input string into a slice of runes,
//	sorting the runes in ascending order, and then converting the sorted runes back into a string.
func SortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
