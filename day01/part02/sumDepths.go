package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// error checking helper function for file reads
func check(e error) {
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

func slidingIncrements(depthsInt []int) int {
	increments := 0
	for i, depths := range depthsInt {
		sumDepths := depths + depthsInt[i+1] + depthsInt[i+2]
		sumDepthsNext := depthsInt[i+1] + depthsInt[i+2] + depthsInt[i+3]

		if sumDepths < sumDepthsNext {
			increments++
		}
	}
	return increments
}

func main() {
	depths, inputErr := getInput("../input.txt")
	check(inputErr)

	depthsInt := strSliceToInt(depths)
	increments := slidingIncrements(depthsInt)

	fmt.Println("Total increments: ", increments)
}
