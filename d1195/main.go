package main

func doesArrayContain(arr []string, val string) bool {
	for _, x := range arr {
		if x == val {
			return true
		}
	}
	return false
}

func getUniqueLetters(input string) []string {
	uniqueLetters := []string{}
	for _, val := range input {
		if !doesArrayContain(uniqueLetters, string(val)) {
			uniqueLetters = append(uniqueLetters, string(val))
		}
	}
	return uniqueLetters
}

func doesSubstringContainAllCharacters(uniqueLetters []string, substr string) bool {
	lettersVisited := make(map[string]int)
	for _, val := range uniqueLetters {
		lettersVisited[val] = 0
	}
	for _, val := range substr {
		lettersVisited[string(val)] += 1
	}
	for key := range lettersVisited {
		if lettersVisited[key] == 0 {
			return false
		}
	}
	return true
}

func getShortestSubstringWithAllInputCharacters(input string) string {
	// get all unique letters in the array
	uniqueLetters := getUniqueLetters(input)

	// if unique letters are the same length as input, return input. there is no substring
	if len(uniqueLetters) == len(input) {
		return input
	}

	// start with the smallest possible substring form index 0 and we will do a growing/sliding window over the string
	// startIndex
	j := 0
	// smallestPossibleEndIndex
	i := len(uniqueLetters)
	currentMinimumSubstring := input

	// x is the tail of the string index that iterates to find a substring
	for x := j + i; x <= len(input); x++ {
		substr := input[j:x]
		if doesSubstringContainAllCharacters(uniqueLetters, substr) && len(substr) <= len(currentMinimumSubstring) {
			currentMinimumSubstring = substr
		}
		j++
	}
	return currentMinimumSubstring
}
