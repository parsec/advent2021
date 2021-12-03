package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) { // error checking helper function for file reads
	if e != nil {
		panic(e)
	}
}

// take path to file, read file line by line into a slice of strings
func getInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// convert []string to []int
func strSliceToInt(inSlice []string) []int {
	intSlice := []int{}
	for _, i := range inSlice {
		tmpInt, atoiErr := strconv.Atoi(i)
		check(atoiErr)
		intSlice = append(intSlice, tmpInt)
	}
	return intSlice
}

func depthIncreases(depths []string) int {
	increments := 0
	// check the depths for increase over the previous
	// skipping the first one of course :)
	for i, depth := range depthInt {
		if i == 0 {
			continue
		} else if depth > depthInt[i-1] {
			increments++
		}
	}

	return increments
}

func main() {
	depths, inputErr := getInput("../input.txt")
	check(inputErr)

	totalIncrements := depthIncreases(depths)

	fmt.Println("The total number of increments is: ", totalIncrements)
}
