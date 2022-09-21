package main

import (
	"log"
	"time"
)

type fn func()

func main() {
	Schedule(SayHello, 5)
}

func SayHello() {
	log.Println("Hello, World!")
}

func Schedule(f fn, n int) {
	second := 1000
	timer := time.NewTimer((time.Millisecond * time.Duration(n*second)))
	<-timer.C
	f()
}
