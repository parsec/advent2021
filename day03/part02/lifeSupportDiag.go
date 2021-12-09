package main

import (
	"fmt"
	"strconv"

	helpers "parsec.sh/lib"
)

// take in slice containing diag data, parse index bit for common bits and remove those that don't match
func oxygenGenRating(diagData []string) int64 {
	mostCommonBits := helpers.CommonBits("most", diagData)
	o2Rating := diagData

	for i := range mostCommonBits {
		o2Rating = helpers.Filter(o2Rating, mostCommonBits[i], i, helpers.CheckBit)
	}

	o2Rate, err := strconv.ParseInt(o2Rating[0], 2, 0)
	helpers.Check(err)

	return o2Rate
}

func cO2ScrubberRating(diagData []string) int64 {
	leastCommonBits := helpers.CommonBits("least", diagData)
	cO2Rating := diagData

	for i := range leastCommonBits {
		cO2Rating = helpers.Filter(cO2Rating, leastCommonBits[i], i, helpers.CheckBit)
	}

	cO2Rate, err := strconv.ParseInt(cO2Rating[0], 2, 0)
	helpers.Check(err)

	return cO2Rate
}

func calcLifeSupport(o2Rate int64, cO2Rate int64) int64 {
	lifeSupportRate := o2Rate * cO2Rate

	return lifeSupportRate
}

func main() {
	diagReport, diagErr := helpers.GetInput("../input.txt")
	helpers.Check(diagErr)

	powerConsumption := calcLifeSupport(oxygenGenRating(diagReport), cO2ScrubberRating(diagReport))

	fmt.Println("Power consumption: ", powerConsumption)
}
