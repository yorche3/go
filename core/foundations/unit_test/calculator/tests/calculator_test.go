package tests

import (
	"testing"

	. "example.com/calculator/src"
	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	result := Addition(2, 3)
	assert.Equal(t, 5, result, "Expected 5 but got %d")
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(5, 2)
	assert.Equal(t, 3, result, "Expected 3 but got %d")
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(4, 3)
	assert.Equal(t, 12, result, "Expected 12 but got %d")
}

func TestDivision(t *testing.T) {
	result := Division(10, 3)
	assert.Equal(t, 3, result, "Expected 3 but got %d")
}

func TestModulus(t *testing.T) {
	result := Modulus(10, 3)
	assert.Equal(t, 1, result, "Expected 1 but got %d")
}
