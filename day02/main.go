package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}(f)

	var reports [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		var report []int
		for _, level := range split {
			converted, err := strconv.Atoi(level)
			if err != nil {
				log.Fatal(err)
			}

			report = append(report, converted)
		}
		reports = append(reports, report)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	safeReports := partOne(reports)
	dampenerReports := partTwo(reports)
	fmt.Printf("%v\n%v\n", safeReports, dampenerReports)
}

func partOne(reports [][]int) int {
	numberSafe := 0

	for _, report := range reports {
		if reportIsSafe(report) {
			numberSafe += 1
		}
	}

	return numberSafe
}

func partTwo(reports [][]int) int {
	numberSafe := 0

	for _, report := range reports {
		works := 0

		for i := 0; i < len(report); i++ {
			smaller := swapRemoveClone(report, i)
			if reportIsSafe(smaller) {
				works += 1
			}
		}

		if works > 0 {
			numberSafe += 1
		}
	}

	return numberSafe
}

func swapRemoveClone(s []int, i int) []int {
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	return append(newSlice[:i], newSlice[i+1:]...)
}

func reportIsSafe(report []int) bool {
	difference := report[1] - report[0]

	for i := 0; i < (len(report) - 1); i++ {
		thisDifference := report[i+1] - report[i]

		goingSameDirection := (difference > 0) == (thisDifference > 0)
		isInRange := (thisDifference >= 1 && thisDifference <= 3) || (thisDifference >= -3 && thisDifference <= -1)

		thisRunWorked := goingSameDirection && isInRange

		if !thisRunWorked {
			return false
		}
	}

	return true
}
