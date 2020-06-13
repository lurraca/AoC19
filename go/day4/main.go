package main

import (
	"github.com/thoas/go-funk"
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
	hash := make(map[string]bool)

	splitInput := strings.Split(input, "")
	for i, char := range splitInput {
		previousCharacter := ""

		if i == len(splitInput)-1 {
			previousCharacter = splitInput[i-1]
			if char == previousCharacter {
				hash[char] = !(char == splitInput[i-2])

				return !(char == splitInput[i-2]) || containsAdjacent(hash)
			}
			return adjacentMatch || containsAdjacent(hash)
		}

		nextCharacter := splitInput[i+1]
		if char != nextCharacter {
			continue
		}

		if adjacentMatch && char == nextCharacter && existOn(hash, char) {
			adjacentMatch = false
			hash[char] = adjacentMatch
			continue
		}

		if char == nextCharacter {
			exists := existOn(hash, char)
			adjacentMatch = !exists
			hash[char] = adjacentMatch && !exists
			continue
		}

		if i != 0 && char == previousCharacter {
			hash[char] = adjacentMatch
			adjacentMatch = false
		}
	}
	return adjacentMatch
}

func existOn(hash map[string]bool, char string) bool {
	return funk.Contains(funk.Keys(hash), char)
}

func containsAdjacent(hash map[string]bool) bool {
	return funk.Some(funk.Values(hash), true)
}
