package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var data []int
	var sum int

	data = readInputs("problem1_2.txt")

	for len(data) > 0 {
		data = calculateFuelRequirements(data)

		for i := 0; i < len(data); i++ {
			sum = sum + data[i]
		}
	}

	fmt.Println(sum)
}

func readInputs(fileName string) []int {
	file, err := os.Open(fileName)

	var data []int

	if err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			mass, _ := strconv.Atoi(scanner.Text())

			data = append(data, mass)
		}

		return data
	}

	return nil
}

func calculateFuelRequirements(modules []int) []int {
	var result []int

	for i := 0; i < len(modules); i++ {
		var fuelReq = calculate(modules[i])

		if fuelReq > 0 {
			result = append(result, fuelReq)
		}
	}

	return result
}

func calculate(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}
