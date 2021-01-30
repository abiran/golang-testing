package services

import (
	"github.com/abiran/golang-testing/src/api/utils/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	elements := sort.GetElements(10)
	Sort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0])
	assert.EqualValues(t, 9, elements[len(elements)-1])
}

func TestSortMoreThan10000(t *testing.T) {
	elements := sort.GetElements(20001)

	Sort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 20001, len(elements))
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 20000, elements[len(elements)-1], "last element should be 20001")
}

func BenchmarkBubbleSort10k(b *testing.B) {
	elements := sort.GetElements(20000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
func BenchmarkBubbleSort100k(b *testing.B) {
	elements := sort.GetElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
