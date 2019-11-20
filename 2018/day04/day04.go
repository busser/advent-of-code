package day04

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"time"
)

// Record represents a record written on the wall.
type Record struct {
	Time time.Time
	Msg  string
}

// GuardID represents the ID of a guard.
type GuardID int

// PartOne solves Part One of the puzzle.
func PartOne(records []Record) int {
	sortRecords(records)

	minutesAsleepByGuard := make(map[GuardID][]int)

	// For each guard, make list of each minute they were asleep.
	var currentGuard GuardID
	var fellAsleepAt time.Time
	for _, rec := range records {
		switch rec.Msg[:5] {
		case "Guard":
			fmt.Sscanf(rec.Msg, "Guard #%d begins shift", &currentGuard)
		case "falls":
			fellAsleepAt = rec.Time
		case "wakes":
			for t := fellAsleepAt; t.Before(rec.Time); t = t.Add(time.Minute) {
				minutesAsleepByGuard[currentGuard] = append(minutesAsleepByGuard[currentGuard], t.Minute())
			}
		}
	}

	// Find guard with most time asleep.
	var sleepiestGuard GuardID
	var mostMinutesAsleep int
	for guard, minutes := range minutesAsleepByGuard {
		if len(minutes) > mostMinutesAsleep {
			sleepiestGuard, mostMinutesAsleep = guard, len(minutes)
		}
	}

	// Find the minute the sleepiest guard sleeps through the most.
	countByMinute := make(map[int]int)
	for _, minute := range minutesAsleepByGuard[sleepiestGuard] {
		countByMinute[minute]++
	}
	var minuteMostSlept, timesSlept int
	for minute, count := range countByMinute {
		if count > timesSlept {
			minuteMostSlept, timesSlept = minute, count
		}
	}

	return int(sleepiestGuard) * minuteMostSlept
}

// PartTwo solves Part Two of the puzzle.
func PartTwo(records []Record) int {
	sortRecords(records)

	minutesAsleepByGuard := make(map[GuardID][]int)

	// For each guard, make list of each minute they were asleep.
	var currentGuard GuardID
	var fellAsleepAt time.Time
	for _, rec := range records {
		switch rec.Msg[:5] {
		case "Guard":
			fmt.Sscanf(rec.Msg, "Guard #%d begins shift", &currentGuard)
		case "falls":
			fellAsleepAt = rec.Time
		case "wakes":
			for t := fellAsleepAt; t.Before(rec.Time); t = t.Add(time.Minute) {
				minutesAsleepByGuard[currentGuard] = append(minutesAsleepByGuard[currentGuard], t.Minute())
			}
		}
	}

	minuteMostSleptByGuard := make(map[GuardID]int)
	timesSleptByGuard := make(map[GuardID]int)

	// For each guard, find the minute they slept through the most.
	for guard, minutesAsleep := range minutesAsleepByGuard {
		countByMinute := make(map[int]int)
		for _, minute := range minutesAsleep {
			countByMinute[minute]++
		}

		var minuteMostSlept, timesSlept int
		for minute, count := range countByMinute {
			if count > timesSlept {
				minuteMostSlept, timesSlept = minute, count
			}
		}

		minuteMostSleptByGuard[guard] = minuteMostSlept
		timesSleptByGuard[guard] = timesSlept
	}

	// Find guard most frequently asleep during the same minute.
	var sleepiestGuard GuardID
	var mostTimesSlept int
	for guard, timesSlept := range timesSleptByGuard {
		if timesSlept > mostTimesSlept {
			sleepiestGuard, mostTimesSlept = guard, timesSlept
		}
	}

	return int(sleepiestGuard) * minuteMostSleptByGuard[sleepiestGuard]
}

// ReadRecords reads records of guards' shifts from r.
func ReadRecords(r io.Reader) ([]Record, error) {
	var records []Record

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		const timeLayout = "[2006-01-02 15:04]"
		time, err := time.Parse(timeLayout, scanner.Text()[:18])
		if err != nil {
			return nil, fmt.Errorf("failed to parse time: %w", err)
		}
		msg := scanner.Text()[19:]

		records = append(records, Record{Time: time, Msg: msg})
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return records, nil
}

// sortRecords sorts records in place.
func sortRecords(records []Record) {
	less := func(i, j int) bool {
		return records[i].Time.Before(records[j].Time)
	}

	if sort.SliceIsSorted(records, less) {
		return
	}

	sort.Slice(records, less)
}
