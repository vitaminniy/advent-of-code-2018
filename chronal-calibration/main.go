package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	frequences := make([]int, 0)

	for scanner.Scan() {
		frequency, _ := strconv.Atoi(scanner.Text())
		frequences = append(frequences, frequency)
	}

	fmt.Printf("%d\n", getResultingFrequency(frequences))
	fmt.Printf("%d\n", getTwiceResultedFrequency(frequences))
}

func getResultingFrequency(frequences []int) int {
	result := 0
	for _, frequency := range frequences {
		result += frequency
	}
	return result
}

func getTwiceResultedFrequency(frequnces []int) int {
	currentFrequency := 0
	calibrations := map[int]int{}

main:
	for {
		for _, frequency := range frequnces {
			currentFrequency += frequency
			if calibrations[currentFrequency]++; calibrations[currentFrequency] == 2 {
				break main
			}
		}
	}

	return currentFrequency
}
