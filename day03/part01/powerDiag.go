package main

import (
	"fmt"
	"strconv"
	"strings"

	helpers "parsec.sh/lib"
)

func calcGammaRate(diagData []string) int {
	dataCompared := helpers.CommonBits("most", diagData)
	rateStr := strings.Join(dataCompared, ",")

	gammaRate, err := strconv.Atoi(rateStr)
	helpers.Check(err)

	return gammaRate
}

func calcEpsilonRate(diagData []string) int {
	dataCompared := helpers.CommonBits("least", diagData)
	rateStr := strings.Join(dataCompared, ",")

	epsilonRate, err := strconv.Atoi(rateStr)
	helpers.Check(err)

	return epsilonRate
}

func main() {
	diagReport, diagErr := helpers.GetInput("../input.txt")
	helpers.Check(diagErr)

	gammaRate := calcGammaRate(diagReport)
	epsilonRate := calcEpsilonRate(diagReport)
	powerConsumption := gammaRate * epsilonRate

	fmt.Println("Power consumption: ", powerConsumption)
}
