package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type entry struct {
	action    string
	timestamp time.Time
}

type entrySlice []entry

func (e entrySlice) Len() int {
	return len(e)
}

func (e entrySlice) Less(i, j int) bool {
	return e[i].timestamp.Before(e[j].timestamp)
}

func (e entrySlice) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func main() {
	filePath := flag.String("p", "input.txt", "Input file path")
	flag.Parse()

	f, err := os.Open(*filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Opening input file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	entries := make(entrySlice, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()

		// [1518-07-08 00:20] wakes up
		var year, month, day, hour, minute int
		if _, err := fmt.Sscanf(text, "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute); err != nil {
			fmt.Fprintf(os.Stderr, "scanning data: %v\n", err)
			os.Exit(1)
		}

		text = text[strings.Index(text, "] ")+2:]

		entries = append(entries, entry{
			action:    text,
			timestamp: time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
		})
	}

	sort.Sort(entries)
	for _, e := range entries {
		fmt.Printf("[%s] %s\n", e.timestamp, e.action)
	}
}
