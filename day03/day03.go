package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func runeInSlice(arr []rune, r rune) bool {
	for _, el := range arr {
		if r == el {
			return true
		}
	}

	return false
}

func getPriority(x rune) int {
	if unicode.IsUpper(x) {
		return int(x) - 38
	} else {
		return int(x) - 96
	}
}

func parseInputP1(input string) []string {
	return strings.Split(input, "\n")
}

func parseInputP2(input string) [][3]string {
	var res [][3]string
	inputSplit := parseInputP1(input)

	for i := 0; i < len(inputSplit); i += 3 {
		var slice [3]string
		copy(slice[:], inputSplit[i:i+3])

		res = append(res, slice)
	}

	return res
}

func getFirstDuplicate(inputs []string) rune {
	for _, letter := range inputs[0] {
		allContain := true

		for _, str := range inputs[1:] {
			if !strings.ContainsRune(str, letter) {
				allContain = false
			}
		}

		if allContain {
			return letter
		}
	}

	return -1
}

func part1(input string) int {
	parsedInput := parseInputP1(input)
	prioritiesSum := 0

	for _, rucksack := range parsedInput {
		c2idx := int(len(rucksack) / 2)
		c1, c2 := rucksack[:c2idx], rucksack[c2idx:]
		duplicate := getFirstDuplicate([]string{c1, c2})

		if duplicate == -1 {
			log.Fatal("No duplicates were found")
		}

		prioritiesSum += getPriority(duplicate)
	}

	return prioritiesSum
}

func part2(input string) int {
	parsedInput := parseInputP2(input)
	prioritiesSum := 0

	for _, group := range parsedInput {
		r1, r2, r3 := group[0], group[1], group[2]
		duplicate := getFirstDuplicate([]string{r1, r2, r3})

		if duplicate == -1 {
			log.Fatal("No duplicates were found")
		}

		prioritiesSum += getPriority(duplicate)
	}

	return prioritiesSum
}

func main() {
	fileData, errRead := os.ReadFile("input.txt")
	dataString := string(fileData)

	if errRead != nil {
		log.Fatal(errRead)
	}

	fmt.Printf("Part 1 : %d\n", part1(dataString))
	fmt.Printf("Part 2 : %d\n", part2(dataString))
}
