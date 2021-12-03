package main

import (
	"fmt"

	helpers "parsec.sh/lib"
)

func depthIncreases(depthInt []int) int {
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
	depths, inputErr := helpers.GetInput("../input.txt")
	helpers.Check(inputErr)

	depthInt := helpers.StrSliceToInt(depths)
	totalIncrements := depthIncreases(depthInt)

	fmt.Println("The total number of increments is: ", totalIncrements)
}
