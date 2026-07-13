package tests

import (
	"testing"

	. "example.com/numbers/src"
	"github.com/stretchr/testify/assert"
)

func TestSumFirstNRec(t *testing.T) {
	assert.Equal(t, 0, SumFirstNRec(0), "Should be equal to 0")
	assert.Equal(t, 6, SumFirstNRec(3), "Should be equal to 6")
}

func TestFactorialRec(t *testing.T) {
	assert.Equal(t, 1, FactorialRec(0), "Should be equal to 1")
	assert.Equal(t, 24, FactorialRec(4), "Should be equal to 24")
}

func TestFibonacciRec(t *testing.T) {
	assert.Equal(t, 0, FibonacciRec(0), "Should be equal to 0")
	assert.Equal(t, 1, FibonacciRec(1), "Should be equal to 1")
	assert.Equal(t, 8, FibonacciRec(6), "Should be equal to 8")
}

func TestGreatestCommonDivisorRec(t *testing.T) {
	assert.Equal(t, 4, GreatestCommonDivisorRec(8, 12), "Should be equal to 4")
	assert.Equal(t, 1, GreatestCommonDivisorRec(7, 5), "Should be equal to 1")
}

func TestLeastCommonMultipleRec(t *testing.T) {
	assert.Equal(t, 24, LeastCommonMultipleRec(8, 6), "Should be 24")
	assert.Equal(t, 12, LeastCommonMultipleRec(6, 4), "Should be 12")
}
