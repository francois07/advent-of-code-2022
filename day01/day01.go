package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) [][]int {
  var res [][]int
  inputSplit := strings.Split(input, "\n\n")

  for _, group := range inputSplit {
    var groupInt []int
    groupSplit := strings.Split(group, "\n")

    for _, x := range groupSplit {
      xInt, errConv := strconv.Atoi(x)

      if errConv != nil {
        log.Fatal(errConv)
      }

      groupInt = append(groupInt, xInt)
    }

    res = append(res, groupInt)
  }

  return res
}

func arraySum(array []int) int {
  res := 0

  for _, x := range array {
    res += x
  }

  return res
}

func part1(input string) ([]int, int) {
  inputInt := parseInput(input)
  var groupCalories []int

  for _, group := range inputInt {
    sum := arraySum(group)
    groupCalories = append(groupCalories, sum)
  }

  sort.Ints(groupCalories)

  return groupCalories, groupCalories[len(groupCalories)-1]
}

func part2(input string) int {
  part1Calories, _ := part1(input)
  topThree := part1Calories[len(part1Calories)-3:]

  return arraySum(topThree)
}

func main() {
  fileData, errRead := os.ReadFile("input.txt")
  dataString := string(fileData)

  if errRead != nil {
    log.Fatal(errRead)
  }

  _, part1Res := part1(dataString)
  part2Res := part2(dataString)

  fmt.Println("Part 1: ", part1Res)
  fmt.Println("Part 2: ", part2Res)
}