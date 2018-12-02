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

	fmt.Printf("%d\n", twices*thirds)

	fmt.Println(findSimilar(ids))
}

func findSimilar(ids []string) string {
	result := ""
main:
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
				fmt.Println(ids[i])
				fmt.Println(ids[j])
				fmt.Println()
				break main
			}
		}
	}
	return result
}
