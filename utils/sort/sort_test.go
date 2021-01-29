package sort

import (
	"testing"
)

func TestBubbleSortOrderDESC(t *testing.T) {
	// Init step
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}
	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}

}

func TestBubbleSortAlreadySorted(t *testing.T) {
	// Init step
	elements := []int{5, 4, 3, 2, 1}

	// Execution
	BubbleSort(elements)

	// Validation ?
}
