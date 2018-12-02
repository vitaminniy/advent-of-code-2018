package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	twices := 0
	thirds := 0

	for scanner.Scan() {
		occurrences := map[rune]int{}

		for _, char := range scanner.Text() {
			occurrences[char]++
		}

		applyTwices := true
		applyThirds := true

		for _, occurrence := range occurrences {
			if occurrence == 2 && applyTwices {
				twices++
				applyTwices = false
			}

			if occurrence == 3 && applyThirds {
				thirds++
				applyThirds = false
			}

			if !applyTwices && !applyThirds {
				break
			}
		}
	}

	fmt.Printf("%d\n", twices*thirds)
}
