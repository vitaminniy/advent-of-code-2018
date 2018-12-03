package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("p", "input.txt", "Path to input file.")
	flag.Parse()

	rectangles := make([]rectangle, 0)

	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "coulnd't open %s: %v\n", *filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rect, err := newRectangle(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't transform input to rectangle: %v\n", err)
			os.Exit(1)
		}

		rectangles = append(rectangles, rect)
	}

	fmt.Printf("Total intersected rectangles: %d\n", countIntersected(rectangles))
}

func countIntersected(rectangles []rectangle) int {
	return 0
}
