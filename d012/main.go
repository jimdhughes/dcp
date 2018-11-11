package main

import "log"

func main() {
	log.Println(StairCombinationsN(4, 3))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func StairCombinations(n int) int {
	return fib(n + 1)
}

func StairCombinationsUtil(n, m int) int {
	if n == 1 {
		return n
	}
	res := 0
	for i := 1; i <= m && i < n; i++ {
		res += StairCombinationsUtil(n-i, m)
	}
	return res
}

func StairCombinationsN(n, m int) int {
	res := StairCombinationsUtil(n+1, m)
	return res
}
