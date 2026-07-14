package src

func SumFirstNRec(n int) int {
	if n <= 0 {
		return 0
	} else {
		return n + SumFirstNRec(n-1)
	}
}

func FactorialRec(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * FactorialRec(n-1)
	}
}

func FibonacciRec(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibonacciRec(n-1) + FibonacciRec(n-2)
	}
}

func GreatestCommonDivisorRec(a, b int) int {
	if b == 0 {
		return a
	} else {
		return GreatestCommonDivisorRec(b, a%b)
	}
}

func LeastCommonMultipleRec(a, b int) int {
	return a * b / GreatestCommonDivisorRec(a, b)
}

func SumFirstNAcc(n int) int {
	return sumFirstHelper(n, 0)
}

func sumFirstHelper(n, acc int) int {
	if n <= 0 {
		return acc
	} else {
		return sumFirstHelper(n-1, acc+n)
	}
}

func FactorialAcc(n int) int {
	return factorialHelper(n, 1)
}

func factorialHelper(n, acc int) int {
	if n <= 1 {
		return acc
	} else {
		return factorialHelper(n-1, acc*n)
	}
}

func FibonacciAcc(n int) int {
	return fibonacciHelper(n, 0, 1)
}

func fibonacciHelper(n, acc1, acc2 int) int {
	if n <= 0 {
		return acc1
	} else {
		return fibonacciHelper(n-1, acc2, acc1+acc2)
	}
}

func SumFirstNIter(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func FactorialIter(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func FibonacciIter(n int) int {
	if n <= 0 {
		return 0
	}
	acc1, acc2 := 0, 1
	for i := 1; i < n; i++ {
		acc1, acc2 = acc2, acc1+acc2
	}
	return acc2
}

func GreatestCommonDivisorIter(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LeastCommonMultipleIter(a, b int) int {
	return (a * b) / GreatestCommonDivisorIter(a, b)
}
