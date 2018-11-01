package main

import "log"

func main() {
	list := []int{3, 2, 1}
	product := 1
	for _, e := range list {
		product *= e
	}
	for i := 0; i < len(list); i++ {
		list[i] = product / list[i]
	}
	log.Println(list)
}
