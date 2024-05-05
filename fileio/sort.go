package fileio

import (
	"bufio"
	"fmt"
	"os"
)

const (
	RadixSort  = "radix"
	MergeSort  = "merge"
	QuickSort  = "quick"
	HeapSort   = "heap"
	RandomSort = "random"
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
	case RadixSort:
		radixSort(lines)
	case MergeSort:
		mergeSort(lines, 0, len(lines)-1)
	case QuickSort:
		quickSort(lines, 0, len(lines)-1)
	case HeapSort:
		heapSort(lines)
	case RandomSort:
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
