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

	calibrations := make(map[int]int)
	currentFrequency := 0
	for {
		for _, frequency := range frequences {
			currentFrequency += frequency
			if calibrations[currentFrequency]++; calibrations[currentFrequency] == 2 {
				fmt.Println(currentFrequency)
				os.Exit(1)
			}
		}
	}
}
