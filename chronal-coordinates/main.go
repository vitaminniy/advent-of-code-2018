package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type coordinate struct {
	x, y int
}

func main() {
	filePath := flag.String("p", "input.txt", "Input's file path")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening input file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	coordinates := make([]coordinate, 0)
	for scanner.Scan() {
		coord := coordinate{}
		if _, err := fmt.Sscanf(scanner.Text(), "%d, %d", &coord.x, &coord.y); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing coordinates: %v\n", err)
			os.Exit(1)
		}
		coordinates = append(coordinates, coord)
	}

	part1(coordinates)
}

func part1(coordinates []coordinate) {

}
