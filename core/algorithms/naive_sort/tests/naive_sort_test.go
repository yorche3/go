package tests

import (
	"testing"

	. "example.com/naive_sort/src"
	"github.com/stretchr/testify/assert"
)

// Casos de prueba de la especificación 05_Naive_Sort.md
var (
	standardInput  = []int{5, 2, 9, 1, 5, 6}
	standardOutput = []int{1, 2, 5, 5, 6, 9}

	sortedInput  = []int{1, 2, 3, 4, 5}
	sortedOutput = []int{1, 2, 3, 4, 5}

	reverseInput  = []int{5, 4, 3, 2, 1}
	reverseOutput = []int{1, 2, 3, 4, 5}

	identicalInput  = []int{7, 7, 7, 7}
	identicalOutput = []int{7, 7, 7, 7}

	negativeInput  = []int{3, -1, 4, -5, 0}
	negativeOutput = []int{-5, -1, 0, 3, 4}

	singleInput  = []int{42}
	singleOutput = []int{42}

	emptyInput  = []int{}
	emptyOutput = []int{}
)

type testCase struct {
	description string
	input       []int
	expected    []int
}

var cases = []testCase{
	{"an unsorted array", standardInput, standardOutput},
	{"an already sorted array", sortedInput, sortedOutput},
	{"a reverse ordered array", reverseInput, reverseOutput},
	{"an array of identical elements", identicalInput, identicalOutput},
	{"an array with negative numbers", negativeInput, negativeOutput},
	{"a single element array", singleInput, singleOutput},
	{"an empty array", emptyInput, emptyOutput},
}

// copyInput evita que un algoritmo in-place contamine los fixtures compartidos.
func copyInput(input []int) []int {
	out := make([]int, len(input))
	copy(out, input)
	return out
}

func assertSortsAllCases(t *testing.T, sort func([]int) []int, algorithm string) {
	t.Helper()

	for _, c := range cases {
		assert.Equal(t, c.expected, sort(copyInput(c.input)),
			algorithm+" should sort "+c.description)
	}

	// Caso nulo: `nil` es el indicador de fallo de Go y se retorna como valor.
	assert.Equal(t, []int(nil), sort(nil),
		algorithm+" should return nil for a nil input")
}

func TestSelectionSort(t *testing.T) {
	assertSortsAllCases(t, SelectionSort, "selection_sort")
}

func TestBubbleSort(t *testing.T) {
	assertSortsAllCases(t, BubbleSort, "bubble_sort")
}

func TestInsertionSort(t *testing.T) {
	assertSortsAllCases(t, InsertionSort, "insertion_sort")
}
