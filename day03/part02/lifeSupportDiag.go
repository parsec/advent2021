package main

import (
	"fmt"
	"strconv"
	"strings"

	helpers "parsec.sh/lib"
)

// take in slice containing diag data, parse index bit for common bits and remove those that don't match
func oxygenGenRating(diagData []string) int64 {
	mostCommonBits := helpers.CommonBits("most", diagData)
	o2RateBits := make([]string, len(mostCommonBits))
	o2Rating := diagData

	for i := range mostCommonBits {
		for x := range o2Rating {
			if x > (len(o2Rating) - 1) { // break if the current index is greater than the current indices
				break
			}
			bits := []rune(o2Rating[x])
			if string(bits[i]) != mostCommonBits[i] {
				o2Rating[x] = o2Rating[len(o2Rating)-1]
				o2Rating[len(o2Rating)-1] = ""
				o2Rating = o2Rating[:len(o2Rating)-1]
			}
		}
	}

	rateStr := strings.Join(o2RateBits, "")
	o2Rate, err := strconv.ParseInt(rateStr, 2, 0)
	helpers.Check(err)

	return o2Rate
}

func cO2ScrubberRating(diagData []string) int64 {
	// todo
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
