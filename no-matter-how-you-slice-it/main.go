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

	fmt.Printf("Total intersected rectangles: %d\n", countIntersectedInches(rectangles))
}

func countIntersectedInches(rectangles []rectangle) int {
	result := 0

	coords := make(map[coordPair]int)
	for _, rect := range rectangles {
		for x := rect.x; x < rect.x+rect.width; x++ {
			for y := rect.y; y < rect.y+rect.height; y++ {
				if coords[coordPair{x, y}]++; coords[coordPair{x, y}] == 2 {
					result++
				}
			}
		}
	}

	return result
}
