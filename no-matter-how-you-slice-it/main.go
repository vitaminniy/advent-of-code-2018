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

	totalOverlapped, notIntactedID := countIntersectedInches(rectangles)
	fmt.Printf("Total intersected rectangles: %d\n", totalOverlapped)
	fmt.Printf("Not intacted rectangle ID: %d\n", notIntactedID)
}

func countIntersectedInches(rectangles []rectangle) (result int, notIntactedID int) {
	coords := make(map[coordPair][]int)
	for _, rect := range rectangles {
		for x := rect.x; x < rect.x+rect.width; x++ {
			for y := rect.y; y < rect.y+rect.height; y++ {
				coords[coordPair{x, y}] = append(coords[coordPair{x, y}], rect.id)
				if len(coords[coordPair{x, y}]) == 2 {
					result++
				}
			}
		}
	}

out:
	for _, rect := range rectangles {
		for x := rect.x; x < rect.x+rect.width; x++ {
			for y := rect.y; y < rect.y+rect.height; y++ {
				if len(coords[coordPair{x, y}]) > 1 {
					continue out
				}
			}
		}
		notIntactedID = rect.id
		break
	}

	return
}
