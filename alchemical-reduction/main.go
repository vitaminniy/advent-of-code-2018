package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {
	filePath := flag.String("p", "input.txt", "Input file path")
	flag.Parse()

	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading input file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Remaining polymer units count: %d\n", len(processPolymer(string(data))))
	fmt.Printf("Optimized polymer units count: %d\n", optimizePolymer(string(data)))
}

func processPolymer(polymer string) string {
	for {
		changes := false
		for k, g := range polymer {
			if k > 0 {
				if unicode.IsLower(g) && unicode.IsUpper(rune(polymer[k-1])) || unicode.IsLower(rune(polymer[k-1])) && unicode.IsUpper(g) {
					if strings.ToLower(string(g)) == strings.ToLower(string(polymer[k-1])) {
						polymer = polymer[:k-1] + polymer[k+1:]
						changes = true
					}
				}
			}
			if changes {
				break
			}
		}
		if !changes {
			break
		}
	}

	return polymer
}

func optimizePolymer(input string) (outcome int) {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"

	outcome = len(input)
	for _, c := range alphabet {
		check := strings.Replace(strings.Replace(input, string(strings.ToUpper(string(c))), "", -1), string(c), "", -1)
		l := len(processPolymer(check))
		if l < outcome {
			outcome = l
		}
	}

	return outcome
}
