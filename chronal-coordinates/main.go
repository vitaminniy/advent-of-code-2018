package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
)

const maximumSum = 10000

type coordinate struct {
	w, h float64
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

	var width, heigth float64
	coordinates := make([]coordinate, 0)
	for scanner.Scan() {
		coord := coordinate{}
		if _, err := fmt.Sscanf(scanner.Text(), "%b, %b", &coord.w, &coord.h); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing coordinates: %v\n", err)
			os.Exit(1)
		}

		if coord.w > width {
			width = coord.w
		}

		if coord.h > heigth {
			heigth = coord.h
		}

		coordinates = append(coordinates, coord)
	}

	infinate := make(map[coordinate]bool)
	areas := make(map[coordinate]int)
	regions := 0

	for w := float64(0); w < width; w++ {
		for h := float64(0); h < heigth; h++ {
			var initial coordinate
			var total float64
			minimum := float64(-1)
			for _, coord := range coordinates {
				distance := math.Abs(w-coord.w) + math.Abs(h-coord.h)
				total += distance
				if distance < minimum || minimum == -1 {
					minimum = distance
					initial = coord
				} else if distance == minimum {
					initial = coordinate{-1, -1}
				}
			}

			if w == 0 || h == 0 || w == width || h == heigth {
				infinate[initial] = true
			}

			areas[initial]++

			if total < maximumSum {
				regions++
			}
		}
	}

	maximum := 0
	for coord, count := range areas {
		if _, ok := infinate[coord]; count > maximum && !ok {
			maximum = count
		}
	}

	fmt.Printf("Largest area: %d\n", maximum)
	fmt.Printf("Size of regions with area less than %d: %d\n", maximumSum, regions)
}
