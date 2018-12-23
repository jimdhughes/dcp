package main

import "log"

func main() {
	log.Println("Divide two positive integers without using division, multiplication, or modulous")
	log.Println(DivideIntegers(7, 8))
}

// DivideIntegers divides the divident by the divisor and returns an integer
// does not care about remainders
func DivideIntegers(dividend, divisor int) int {
	iteration := 0
	res := 0
	if divisor == 0 {
		return 0
	}
	for res < dividend {
		if res+divisor > dividend {
			break
		}
		res += divisor
		iteration++
	}
	return iteration
}
