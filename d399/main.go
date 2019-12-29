package main

import "log"

func sumPartition(partition []int) int {
	t := 0
	for _, v := range partition {
		t += v
	}
	return t
}

func findThreePartitions(partition []int) [][]int {
	firstPartitionIdx := 1
	secondPartitionIdx := 2
	var validPartitions [][]int
	for firstPartitionIdx < len(partition) && secondPartitionIdx < len(partition) {
		p1 := partition[:firstPartitionIdx]
		p2 := partition[firstPartitionIdx:secondPartitionIdx]
		p3 := partition[secondPartitionIdx:]
		if sumPartition(p1) == sumPartition(p2) && sumPartition(p2) == sumPartition(p3) {
			validPartitions = [][]int{p1, p2, p3}
			return validPartitions
		}
		secondPartitionIdx++
		if secondPartitionIdx == len(partition) && firstPartitionIdx < len(partition) {
			firstPartitionIdx++
			secondPartitionIdx = firstPartitionIdx + 1
		}
	}
	return validPartitions
}

func main() {
	log.Println(findThreePartitions([]int{3, 5, 8, 0, 8}))
}
