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
	timer := time.NewTimer((time.Millisecond * 2))
	<-timer.C
	f()
}
