package main

import (
	"fmt"
	"strconv"
	"strings"

	helpers "parsec.sh/lib"
)

// takes []string of commands and returns x, y values
func findPosition(commands []string) (int, int) {
	x, y := 0, 0

	for _, cmd := range commands {
		splitCmd := strings.Split(cmd, " ")
		direction := splitCmd[0]
		amount, atoiErr := strconv.Atoi(splitCmd[1])
		helpers.Check(atoiErr)

		switch direction {
		case "up":
			y -= amount
		case "down":
			y += amount
		case "forward":
			x += amount
		}
	}
	return x, y
}

func main() {
	input, inputErr := helpers.GetInput("../inputCommands.txt")
	helpers.Check(inputErr)

	xPos, yPos := findPosition(input)
	product := xPos * yPos

	fmt.Println("X: ", xPos)
	fmt.Println("Y: ", yPos)
	fmt.Println("Product of X*Y: ", product)
}
