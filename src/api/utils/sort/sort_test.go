package sort

import (
	"testing"
)

func TestBubbleSortOrderIncreasing(t *testing.T) {
	// Init step
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}

}

func TestSortOrderIncreasing(t *testing.T) {
	// Init step
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution
	Sort(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}
	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}

}

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}
func BenchmarkSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
