package helpers

import (
	"bufio"
	"os"
	"strconv"
)

// error checking helper function for file reads
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// convert []string to []int
func StrSliceToInt(inSlice []string) []int {
	intSlice := []int{}
	for _, i := range inSlice {
		tmpInt, atoiErr := strconv.Atoi(i)
		Check(atoiErr)
		intSlice = append(intSlice, tmpInt)
	}
	return intSlice
}

// take path to file, read file line by line into a slice of strings
func GetInput(path string) ([]string, error) {
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
