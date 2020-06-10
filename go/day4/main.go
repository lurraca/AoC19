package main

import (
	"strconv"
	"strings"
)

func main() {
	count := 0
	for i := 136818; i < 685979; i++ {
		if isPasswordCandidate(strconv.Itoa(i)) {
			count = count + 1
		}
	}
	println(count)
}

func isPasswordCandidate(input string) bool {
	if !isAdjacentMatch(input) {
		return false
	}
	if decreases(input) {
		return false
	}
	return true
}

func decreases(input string) bool {
	var decreases bool
	splitInput := strings.Split(input, "")
	for i, char := range splitInput {
		if decreases {
			return true
		}
		currentCharInt, _ := strconv.Atoi(char)
		if i == len(splitInput)-1 {
			previousCharInt, _ := strconv.Atoi(splitInput[i-1])
			return currentCharInt < previousCharInt
		}
		nextCharInt, _ := strconv.Atoi(splitInput[i+1])
		if currentCharInt > nextCharInt {
			decreases = true
		}
	}
	return decreases
}

func isAdjacentMatch(input string) bool {
	var adjacentMatch bool
	splitInput := strings.Split(input, "")
	for i, char := range splitInput {
		if adjacentMatch {
			return true
		}
		if i == len(splitInput)-1 {
			return char == splitInput[i-i]
		}
		if char == splitInput[i+1] {
			adjacentMatch = true
		}
	}
	return adjacentMatch
}
