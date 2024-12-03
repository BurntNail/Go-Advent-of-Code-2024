package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
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

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	multiplySum := partOne(lines)
	betterMultiplySum := partTwo(lines)
	fmt.Printf("%v\n%v\n", multiplySum, betterMultiplySum)
}

func partOne(lines []string) int {
	sum := 0
	isEnabled := true
	for _, line := range lines {
		sum += processLine(line, true, &isEnabled)
	}
	return sum
}

func partTwo(lines []string) int {
	sum := 0
	isEnabled := true
	for _, line := range lines {
		sum += processLine(line, false, &isEnabled)
	}
	return sum
}

func processLine(line string, ignoreDisable bool, isEnabled *bool) int {
	if isEnabled == nil {
		return -1
	}

	runes := []rune(line)
	sum := 0

	var leftNumber int
	var rightNumber int
	tmp := ""
	state := 0

	for i := 0; i < len(runes); i++ {
		fmt.Printf("tmp: \"%s\", ch: %c, state: %d, left: %v, right: %v\n", tmp, runes[i], state, leftNumber, rightNumber)

		switch state {
		default:
			log.Fatal("unreachable")
		case 0:
			tmp = tmp + string(runes[i])

			if *isEnabled || ignoreDisable {
				switch len(tmp) {
				case 1:
					if tmp != "m" && tmp != "d" {
						tmp = ""
					}
				case 2:
					if tmp != "mu" && tmp != "do" {
						tmp = ""
					}
				case 3:
					if tmp != "mul" && tmp != "don" {
						tmp = ""
					}
				case 4:
					if tmp != "mul(" && tmp != "don'" {
						tmp = ""
					}

					if tmp == "mul(" {
						state = 1
						tmp = ""
					}
				case 5:
					if tmp != "don't" {
						tmp = ""
					}
				case 6:
					if tmp != "don't(" {
						tmp = ""
					}
				case 7:
					if tmp == "don't()" {
						*isEnabled = false
					}
					tmp = ""
				}
			} else {
				switch len(tmp) {
				case 1:
					if tmp != "d" {
						tmp = ""
					}
				case 2:
					if tmp != "do" {
						tmp = ""
					}
				case 3:
					if tmp != "do(" {
						tmp = ""
					}
				case 4:
					if tmp == "do()" {
						*isEnabled = true
					}
					tmp = ""
				}
			}
		case 1:
			if unicode.IsDigit(runes[i]) {
				tmp = tmp + string(runes[i])
				continue
			}

			if runes[i] != ',' {
				tmp = ""
				state = 0
				continue
			}

			converted, err := strconv.Atoi(tmp)
			if err != nil {
				tmp = ""
				state = 0
				continue
			}

			if converted > 999 {
				tmp = ""
				state = 0
				continue
			}

			leftNumber = converted
			tmp = ""
			state = 2
		case 2:
			if unicode.IsDigit(runes[i]) {
				tmp = tmp + string(runes[i])
				continue
			}

			if runes[i] != ')' {
				tmp = ""
				state = 0
				continue
			}

			converted, err := strconv.Atoi(tmp)
			if err != nil {
				tmp = ""
				state = 0
				continue
			}

			if converted > 999 {
				tmp = ""
				state = 0
				continue
			}

			rightNumber = converted
			sum += leftNumber * rightNumber

			tmp = ""
			state = 0
		}
	}

	return sum
}
