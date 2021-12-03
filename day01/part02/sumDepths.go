package main

import (
	"fmt"

	helpers "parsec.sh/lib"
)

func slidingIncrements(depthsInt []int) int {
	increments := 0
	for i, depths := range depthsInt {
		if i == (len(depthsInt) - 3) { // break if we've reached the end of the slice
			break
		}

		sumDepths := depths + depthsInt[i+1] + depthsInt[i+2]
		sumDepthsNext := depthsInt[i+1] + depthsInt[i+2] + depthsInt[i+3] // there's probably a better way to do this that doesn't require this janky stuff

		if sumDepths < sumDepthsNext {
			increments++
		}
	}
	return increments
}

func main() {
	depths, inputErr := helpers.GetInput("../input.txt")
	helpers.Check(inputErr)

	depthsInt := helpers.StrSliceToInt(depths)
	increments := slidingIncrements(depthsInt)

	fmt.Println("Total increments: ", increments)
}
