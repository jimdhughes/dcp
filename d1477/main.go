package main

import "log"

func main() {
	log.Println(GetFirstNRegularNumbers(60))
}

func GetFirstNRegularNumbers(n int) []int {
	regularNumbers := []int{1}
	// get all prime numbers in range of 1 to n
	primes := GetAllPrimeNumbers(n)
	// get the first 3 prime numbers
	validPrimes := primes[0:2]
	// remove the first 3 prime numbers from the list
	primes = primes[3:]
	// get all regular numbers in range of 1 to n
	for i := 1; i <= n; i++ {
		valid := false
		// check if the number is divisible by the first 3 prime numbers
		for _, validPrime := range validPrimes {
			if i%validPrime == 0 {
				valid = true
			}
		}
		// check if the number is divisible by the rest of the prime numbers
		for _, primes := range primes {
			if i%primes == 0 {
				// if the number is divisible by any of the prime numbers, it is not a regular number
				valid = false
			}
		}
		// if the number is not divisible by any of the prime numbers, it is a regular number
		if valid {
			regularNumbers = append(regularNumbers, i)
		}
	}
	return regularNumbers
}

func GetAllPrimeNumbers(input int) []int {
	var primeNumbers []int
	for i := 1; i <= input; i++ {
		if IsPrimeNumber(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}

func IsPrimeNumber(input int) bool {
	if input < 2 {
		return false
	}
	for i := 2; i < input; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}
