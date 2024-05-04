package fileio

import (
	"bufio"
	"fmt"
	"os"
)

func SortFile(file *os.File, unique bool, sortAlgo string) []string {

	// Read lines from file
	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Sort the lines
	switch sortAlgo {
	case "radix":
		radixSort(lines)
	case "merge":
		mergeSort(lines, 0, len(lines)-1)
	case "quick":
		quickSort(lines, 0, len(lines)-1)
	case "heap":
		heapSort(lines)
	case "random":
		randomSort(lines)
	default:
		fmt.Println("Invalid sort algorithm specified")
		os.Exit(1)
	}

	// Remove duplicates
	if !unique {
		return lines
	}

	uniqueLines := removeDuplicates(lines)

	return uniqueLines

}

func removeDuplicates(lines []string) []string {
	uniqueLines := make(map[string]bool)
	var result []string

	for _, line := range lines {
		if !uniqueLines[line] {
			uniqueLines[line] = true
			result = append(result, line)
		}
	}

	return result
}
