package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

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

	twices := 0
	thirds := 0

	ids := []string{}

	for scanner.Scan() {
		occurrences := map[rune]int{}

		txt := scanner.Text()

		for _, char := range txt {
			occurrences[char]++
		}

		ids = append(ids, txt)

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

	fmt.Printf("Checksum: %d\n", twices*thirds)
	fmt.Printf("Common letters: %s\n", findCommonLetters(ids))
}

func findCommonLetters(ids []string) (result string) {
	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(ids); j++ {
			if i == j {
				continue
			}

			diffCount := 0

			for idx := range ids[i] {
				if ids[i][idx] != ids[j][idx] {
					diffCount++
				}

				if diffCount > 1 {
					break
				}
			}

			if diffCount == 1 {
				for idx := range ids[i] {
					if ids[i][idx] == ids[j][idx] {
						result += string(ids[i][idx])
					}
				}
				return result
			}
		}
	}
	return
}
