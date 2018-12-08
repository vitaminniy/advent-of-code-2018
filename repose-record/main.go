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

type action int

const (
	beginShift action = iota
	fallAsleep
	wakeUp
)

func (a *action) String() string {
	return [...]string{"begins shift", "falls asleep", "wakes up"}[*a]
}

type entry struct {
	guard     int
	action    action
	timestamp time.Time
}

func (e *entry) String() string {
	if e.action == beginShift {
		return fmt.Sprintf("[%s] Guard #%d %s", e.timestamp, e.guard, e.action.String())
	}

	return fmt.Sprintf("[%s] %s", e.timestamp, e.action.String())
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

		e := entry{timestamp: time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)}
		text = text[strings.Index(text, "] ")+2:]
		n, _ := fmt.Sscanf(text, "Guard #%d begins shift", &e.guard)

		switch {
		case n == 1:
			e.action = beginShift
		case text == "falls asleep":
			e.action = fallAsleep
		case text == "wakes up":
			e.action = wakeUp
		}

		entries = append(entries, e)
	}

	sort.Sort(entries)

	firstStrategy(entries)
	secondStrategy(entries)
}

func firstStrategy(entries []entry) {
	var sleepyguard int
	asleep := map[int]int{}
	var guard, from int
	for _, e := range entries {
		switch e.action {
		case beginShift:
			guard = e.guard
		case fallAsleep:
			from = e.timestamp.Minute()
		case wakeUp:
			t := e.timestamp.Minute() - from
			asleep[guard] += t
			if asleep[guard] > asleep[sleepyguard] {
				sleepyguard = guard
			}
		}
	}

	minutes := [60]int{}
	guard = -1
	var sleepyminute int
	for _, e := range entries {
		if e.action == beginShift {
			guard = e.guard
			continue
		}
		if guard != sleepyguard {
			continue
		}
		switch e.action {
		case fallAsleep:
			from = e.timestamp.Minute()
		case wakeUp:
			to := e.timestamp.Minute()
			for i := from; i < to; i++ {
				minutes[i]++
				if minutes[i] > minutes[sleepyminute] {
					sleepyminute = i
				}
			}
		}
	}

	fmt.Printf("Answer: guard %d * minute %d = %d\n",
		sleepyguard, sleepyminute, sleepyguard*sleepyminute)
}

func secondStrategy(entries []entry) {
	var sleepyguard, sleepyminute int
	minutes := map[int]*[60]int{}
	var guard, from int
	for _, e := range entries {
		switch e.action {
		case beginShift:
			guard = e.guard
			if minutes[guard] == nil {
				minutes[guard] = &[60]int{}
			}
			if minutes[sleepyguard] == nil {
				sleepyguard = guard
			}

		case fallAsleep:
			from = e.timestamp.Minute()
		case wakeUp:
			to := e.timestamp.Minute()
			for i := from; i < to; i++ {
				minutes[guard][i]++
				if minutes[guard][i] > minutes[sleepyguard][sleepyminute] {
					sleepyguard = guard
					sleepyminute = i
				}
			}
		}
	}

	fmt.Printf("Answer: guard %d * minute %d = %d\n",
		sleepyguard, sleepyminute, sleepyguard*sleepyminute)
}
