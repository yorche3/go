package src

func Addition(a, b int) int {
	return a + b
}

func Subtraction(a, b int) int {
	return a - b
}

func Multiplication(a, b int) int {
	result := 0
	for i := 0; i < b; i++ {
		result = Addition(result, a)
	}
	return result
}

func Division(a, b int) int {
	quotient := 0
	for a >= b {
		a = Subtraction(a, b)
		quotient = Addition(quotient, 1)
	}
	return quotient
}

func Modulus(a, b int) int {
	quotient := Division(a, b)
	return Subtraction(a, Multiplication(b, quotient))
}
