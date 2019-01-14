package main

import (
	"log"
	"strconv"
)

const (
	odd  = 0xAA // 10101010
	even = 0x55 // 01010101
)

func main() {
	log.Println("Swapping some binary")
	log.Printf("Swapping 5 - 0101: %s", PrintBinary(SwapBinary(5)))
	log.Printf("Swapping 10 - 1010: %s", PrintBinary(SwapBinary(10)))
	log.Printf("Swapping 226 - 11100010: %s", PrintBinary(SwapBinary(226)))
}

func PrintBinary(number uint8) string {
	return strconv.FormatInt(int64(number), 2)
}

func SwapBinary(number uint8) uint8 {
	return (number & even << 1) | (number & odd >> 1)
}
