package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) PopN(n int) ([]string, bool) {
	if len(*s) < n {
		return nil, false
	} else {
		var res []string

		for i := 0; i < n; i++ {
			index := len(*s) - n + i
			element := (*s)[index]
			res = append(res, element)
		}
		*s = (*s)[:len(*s)-n]

		return res, true
	}
}

func (s *Stack) PushN(strs []string) {
	for _, str := range strs {
		*s = append(*s, str)
	}
}

func parseInput(input string) ([]Stack, [][3]int) {
	inputSplit := strings.Split(input, "\n\n")
	rawStacks := strings.Split(inputSplit[0], "\n")
	rawInstructions := strings.Split(inputSplit[1], "\n")

	var stacks []Stack
	var instructions [][3]int

	/* Parse crate stacks */
	for i := 1; i < len(rawStacks[0]); i += 4 {
		var newStack Stack

		for j := len(rawStacks) - 2; j >= 0; j-- {
			letter := rawStacks[j][i]

			if letter != ' ' {
				newStack.Push(string(letter))
			}
		}

		stacks = append(stacks, newStack)
	}

	/* Parse instructions */
	r, _ := regexp.Compile("[0-9]+")

	for _, line := range rawInstructions {
		matches := r.FindAllString(line, 3)
		var instruction [3]int

		for i, m := range matches {
			if mInt, err := strconv.Atoi(m); err == nil {
				instruction[i] = mInt
			}
		}

		instructions = append(instructions, instruction)
	}

	return stacks, instructions
}

func getTopOfStacks(stacks []Stack) string {
	var res string

	for _, stack := range stacks {
		if !stack.IsEmpty() {
			res += stack[len(stack)-1]
		}
	}

	return res
}

func part1(input string) string {
	parsedStacks, parsedInstructions := parseInput(input)
	finalStacks := parsedStacks

	/* Execute instructions */
	for _, inst := range parsedInstructions {
		quantity, source, destination := inst[0], inst[1]-1, inst[2]-1

		for i := 0; i < quantity; i++ {
			if popped, success := finalStacks[source].Pop(); success {
				finalStacks[destination].Push(popped)
			}
		}
	}

	return getTopOfStacks(finalStacks)
}

func part2(input string) string {
	parsedStacks, parsedInstructions := parseInput(input)
	finalStacks := parsedStacks

	/* Execute instructions */
	for _, inst := range parsedInstructions {
		quantity, source, destination := inst[0], inst[1]-1, inst[2]-1

		if popped, success := finalStacks[source].PopN(quantity); success {
			finalStacks[destination].PushN(popped)
		}
	}

	return getTopOfStacks(finalStacks)
}

func main() {
	fileData, errRead := os.ReadFile("input.txt")
	dataString := string(fileData)

	if errRead != nil {
		log.Fatal(errRead)
	}

	fmt.Printf("Part 1 : %s\n", part1(dataString))
	fmt.Printf("Part 2 : %s\n", part2(dataString))
}
