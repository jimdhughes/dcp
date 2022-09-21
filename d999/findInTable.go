package main

func CountOccurancesInMatrix(matrixSize, value int) int {
	i := 1
	j := matrixSize
	counts := 0

	for i <= j {
		if i*j == value {
			if i == j {
				// this is the diagonal
				counts += 1
			} else {
				// this is for above / below diagnoal
				counts += 2
			}
			i++
			continue
		}
		if i*j < value {
			i++
		}
		if i*j > value {
			j--
		}
	}

	return counts
}

// create a function that finds the number of times a number appears in a multiplication table
func GetNumberOfOccurancesInMultiplicationTable(matrixSize, value int) int {
	var counts int
	for i := 1; i <= matrixSize; i++ {
		for j := 1; j <= matrixSize; j++ {
			if i*j == value {
				counts++
			}
		}
	}
	return counts
}
