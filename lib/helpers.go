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

// convert an []int to a []string
func IntToStrSlice(intSlice []int) []string {
	strSlice := []string{}

	for i := range intSlice {
		num := strconv.Itoa(intSlice[i])
		strSlice = append(strSlice, num)
	}

	return strSlice
}

func CommonBits(howCommon string, data []string) []string {
	zeroes := make([]int, len(data[0]))
	ones := make([]int, len(data[0]))
	results := make([]int, len(data[0]))

	// iterate over the data, line by line (bit by bit!) to count bit occurences
	for i := range data {
		bits := []rune(data[i])
		for x := range bits {
			if string(bits[x]) == "0" {
				zeroes[x] += 1
				ones[x] += 0
			} else if string(bits[x]) == "1" {
				ones[x] += 1
				zeroes[x] += 0
			}
		}
	}

	// compare both int slices to determine which value was greater for that column, 0 or 1
	// in this use case, we want to be able to get either least or most common bits for specific values
	switch howCommon {
	case "least":
		for i := range ones {
			if ones[i] < zeroes[i] {
				results[i] = 1
			} else if ones[i] > zeroes[i] {
				results[i] = 0
			}
		}
	case "most":
		for i := range ones {
			if ones[i] > zeroes[i] {
				results[i] = 1
			} else if ones[i] < zeroes[i] {
				results[i] = 0
			}
		}
	default:
		results[0] = -1
	}

	resultsStr := IntToStrSlice(results)
	return resultsStr
}
