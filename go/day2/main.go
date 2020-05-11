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
	input := readInput()
	input[1] = 12
	input[2] = 2
	fmt.Println(Intcode(input))

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
