package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println(calculateMass("problem1_1.txt"))
}

func calculateMass(fileName string) int {
	file, err := os.Open(fileName)

	var sum = 0

	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			mass, _ := strconv.Atoi(scanner.Text())

			sum = sum + calculateFuelRequirements(mass)
		}

		return sum
	}

	return 0
}

func calculateFuelRequirements(moduleMass int) int {
	return int(math.Floor(float64(moduleMass)/3)) - 2
}
