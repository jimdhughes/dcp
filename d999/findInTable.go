package main

func FindInMatrix(matrixSize, value int) int {
	i := 1
	j := matrixSize
	counts := 0

	for i <= j {
		if i*j == value {
			if (i == matrixSize && j == matrixSize) || (i == 1 && j == 1) {
				counts += 1
			} else {
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
