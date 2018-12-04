package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filePath := flag.String("p", "input.txt", "Input file")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening inout file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	frequences := make([]int, 0)

	for scanner.Scan() {
		frequency, _ := strconv.Atoi(scanner.Text())
		frequences = append(frequences, frequency)
	}

	fmt.Printf("Resulting frequency: %d\n", getResultingFrequency(frequences))
	fmt.Printf("Resulting frequency reached twice: %d\n", getTwiceResultedFrequency(frequences))
}

func getResultingFrequency(frequences []int) int {
	result := 0
	for _, frequency := range frequences {
		result += frequency
	}
	return result
}

func getTwiceResultedFrequency(frequnces []int) (result int) {
	calibrations := map[int]int{}

	for {
		for _, frequency := range frequnces {
			result += frequency
			if calibrations[result]++; calibrations[result] == 2 {
				return
			}
		}
	}
}
