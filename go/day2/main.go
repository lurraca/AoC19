//--- Day 2: 1202 Program Alarm ---
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	partA()
	partB()
}

func partA() {
	fmt.Printf("Part A: %d\n", IntcodeArgs(12, 2))
}

func partB() {
	target := 19690720

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if target == IntcodeArgs(noun, verb)[0] {
				fmt.Printf("Part B: Noun is: %d, verb is %d an answer is %v ", noun, verb, 100*noun+verb)
			}
		}
	}
}

func IntcodeArgs(noun int, verb int) []int {
	input := readInput()
	input[1] = noun
	input[2] = verb
	return Intcode(input)
}

func Intcode(input []int) []int {
	batches := batches(4, input)

	for _, batch := range batches {
		opcode := batch[0]
		switch {
		case opcode == 1:
			input[batch[3]] = input[batch[1]] + input[batch[2]]
		case opcode == 2:
			input[batch[3]] = input[batch[1]] * input[batch[2]]
		case opcode == 99:
			return input
		}
	}

	return input
}

func batches(batchSize int, input []int) [][]int {
	batches := make([][]int, 0, (len(input)+batchSize-1)/batchSize)

	for batchSize < len(input) {
		input, batches = input[batchSize:], append(batches, input[0:batchSize:batchSize])
	}
	batches = append(batches, input)
	return batches
}

func readInput() []int {
	content, err := ioutil.ReadFile("day2\\input.txt")
	if err != nil {
		log.Fatal(err)
	}

	collectionIntAsStrings := strings.Split(string(content), ",")
	var collectionInt []int
	for _, e := range collectionIntAsStrings {
		eAsInt, _ := strconv.Atoi(e)
		collectionInt = append(collectionInt, eAsInt)
	}
	return collectionInt
}
