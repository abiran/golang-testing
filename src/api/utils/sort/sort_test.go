package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortOrderIncreasing(t *testing.T) {
	elements := getElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
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
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}
func BenchmarkSort(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
