package main

import (
	"fmt"

	helpers "parsec.sh/lib"
)

func calcGammaRate(diagData []string) int {

}

func calcEpsilonRate(diagData []string) int {

}

func main() {
	diagReport, diagErr := helpers.GetInput("../input.txt")
	helpers.Check(diagErr)

	gammaRate := calcGammaRate(diagReport)
	epsilonRate := calcEpsilonRate(diagReport)
	powerConsumption := gammaRate * epsilonRate

	fmt.Println("Power consumption: ", powerConsumption)
}
