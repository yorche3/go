package tests

import (
	"testing"

	. "example.com/numbers/src"
	"github.com/stretchr/testify/assert"
)

func TestSumFirstNIter(t *testing.T) {
	assert.Equal(t, 0, SumFirstNIter(0), "Should be equal to 0")
	assert.Equal(t, 6, SumFirstNIter(3), "Should be equal to 6")
}

func TestFactorialIter(t *testing.T) {
	assert.Equal(t, 1, FactorialIter(0), "Should be equal to 1")
	assert.Equal(t, 24, FactorialIter(4), "Should be equal to 24")
}

func TestFibonacciIter(t *testing.T) {
	assert.Equal(t, 0, FibonacciIter(0), "Should be equal to 0")
	assert.Equal(t, 1, FibonacciIter(1), "Should be equal to 1")
	assert.Equal(t, 8, FibonacciIter(6), "Should be equal to 8")
}

func TestGreatestCommonDivisorIter(t *testing.T) {
	assert.Equal(t, 4, GreatestCommonDivisorIter(12, 8), "Should be equal to 4")
	assert.Equal(t, 1, GreatestCommonDivisorIter(7, 5), "Should be equal to 1")
}

func TestLeastCommonMultipleIter(t *testing.T) {
	assert.Equal(t, 24, LeastCommonMultipleIter(6, 8), "Should be equal to 24")
	assert.Equal(t, 12, LeastCommonMultipleIter(6, 4), "Should be equal to 12")
}
