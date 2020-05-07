//--- Day 1: The Tyranny of the Rocket Equation ---
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	totalFuel := 0
	for _, inputMass := range readInput() {
		mass, _ := strconv.Atoi(inputMass)
		totalFuel = totalFuel + Fuel(mass)
	}

	fmt.Printf("Total fuel is: %d", totalFuel)
}

func readInput() []string {
	var input []string
	file, err := os.Open("day1\\input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func Fuel(mass int) int {
	return (mass / 3) - 2
}
