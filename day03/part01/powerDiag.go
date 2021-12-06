package main

import (
	"fmt"
	"strconv"
	"strings"

	helpers "parsec.sh/lib"
)

func calcGammaRate(diagData []string) int64 {
	dataCompared := helpers.CommonBits("most", diagData)
	rateStr := strings.Join(dataCompared, "")

	gammaRate, err := strconv.ParseInt(rateStr, 2, 0)
	helpers.Check(err)

	return gammaRate
}

func calcEpsilonRate(diagData []string) int64 {
	dataCompared := helpers.CommonBits("least", diagData)
	rateStr := strings.Join(dataCompared, "")

	epsilonRate, err := strconv.ParseInt(rateStr, 2, 0)
	helpers.Check(err)

	return epsilonRate
}

func calcPowerConsumption(gammaRate int64, epsilonRate int64) int64 {

	powerConsumption := gammaRate * epsilonRate

	return powerConsumption
}

func main() {
	diagReport, diagErr := helpers.GetInput("../input.txt")
	helpers.Check(diagErr)

	powerConsumption := calcPowerConsumption(calcGammaRate(diagReport), calcEpsilonRate(diagReport))

	fmt.Println("Power consumption: ", powerConsumption)
}
