package main

import "testing"

func TestArraySum(t *testing.T) {
	expectedSum := 45
	nums := [5]int{5, 10, 15, 5, 10}

	result := SumArray(nums)
	if result != expectedSum {
		t.Errorf("Sum of Array nums Wrong! expected %d, got %d; input %v", expectedSum, result, nums)
	}
}

func TestSliceSum(t *testing.T) {
	expectedSum := 21
	nums := []int{5, 3, 5, 4, 4}

	result := SumSlice(nums)
	if result != expectedSum {
		t.Errorf("Sum of Slice nums Wrong! expected %d, got %d; input %v", expectedSum, result, nums)
	}

}
