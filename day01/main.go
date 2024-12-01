package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	var left, right []int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")

		leftNumber, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, leftNumber)

		rightNumber, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, rightNumber)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//sumDistances := partOne(left, right)
	//fmt.Printf("%d\n", sumDistances)

	similarity := partTwo(left, right)
	fmt.Printf("%d\n", similarity)
}

func partOne(left []int, right []int) int {
	var distances []int
	for _ = range len(left) {
		leftSmallest, newLeft := getSmallestAndRemove(left)
		rightSmallest, newRight := getSmallestAndRemove(right)

		left = newLeft
		right = newRight

		if leftSmallest > rightSmallest {
			distances = append(distances, leftSmallest-rightSmallest)
		} else {
			distances = append(distances, rightSmallest-leftSmallest)
		}
	}

	sumDistances := 0
	for _, dist := range distances {
		sumDistances += dist
	}

	return sumDistances
}

func getSmallestAndRemove(list []int) (int, []int) {
	if len(list) == 1 {
		return list[0], []int{}
	}

	smallest := math.MaxInt
	smallestIndex := 0

	for index, element := range list {
		if element < smallest {
			smallest = element
			smallestIndex = index
		}
	}

	newList := swapRemove(list, smallestIndex)

	return smallest, newList
}

func swapRemove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func partTwo(left []int, right []int) int {
	var similarities []int

	for _, item := range left {
		score := countOccurences(right, item)
		similarities = append(similarities, item*score)
	}

	similarityTotal := 0
	for _, sim := range similarities {
		similarityTotal += sim
	}

	return similarityTotal
}

func countOccurences(list []int, element int) int {
	count := 0
	for _, item := range list {
		if item == element {
			count += 1
		}
	}
	return count
}
