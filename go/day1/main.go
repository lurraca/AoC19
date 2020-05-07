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
	recursiveTotalFuel := 0
	for _, inputMass := range readInput() {
		mass, _ := strconv.Atoi(inputMass)
		totalFuel = totalFuel + Fuel(mass)
		recursiveTotalFuel = recursiveTotalFuel + RecursiveFuel(mass)
	}

	fmt.Printf("Total fuel is: %d \n", totalFuel)
	fmt.Printf("Total recursive fuel is: %d", recursiveTotalFuel)
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

func RecursiveFuel(mass int) int {
	if Fuel(mass) <= 0 {
		return 0
	} else {
		return Fuel(mass) + RecursiveFuel(Fuel(mass))
	}
}
