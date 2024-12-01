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

	sumDistances := partOne(left, right)
	fmt.Printf("%d\n", sumDistances)

	similarity := partTwo(left, right)
	fmt.Printf("%d\n", similarity)
}

func partOne(remoteLeft []int, remoteRight []int) int {
	left := make([]int, len(remoteLeft))
	right := make([]int, len(remoteRight))

	copy(left, remoteLeft)
	copy(right, remoteRight)

	distance := 0
	for range len(left) {
		leftSmallest, newLeft := getSmallestAndRemove(left)
		rightSmallest, newRight := getSmallestAndRemove(right)

		left = newLeft
		right = newRight

		if leftSmallest > rightSmallest {
			distance += leftSmallest - rightSmallest
		} else {
			distance += rightSmallest - leftSmallest
		}
	}

	return distance
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
	similarity := 0

	var occurences = make(map[int]int)

	for _, item := range right {
		occurences[item] += 1
	}

	for _, item := range left {
		similarity += item * occurences[item]
	}

	return similarity
}
