package main

import "log"

func getPerm(arr []string, perm []int) []string {
	res := []string{}
	for _, v := range perm {
		res = append(res, arr[v])
	}
	return res
}

func main() {
	arr := []string{"a", "b", "c"}
	perm := []int{2, 1, 0}
	log.Println(getPerm(arr, perm))
}
