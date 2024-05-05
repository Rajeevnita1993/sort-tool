package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Rajeevnita1993/sort-tool/fileio"
)

func main() {

	sortAlgo := flag.String("algo", fileio.QuickSort, "Select sorting algorithm: radix, merge, quick, heap")
	unique := flag.Bool("u", false, "Remove dupplicate lines")
	flag.Parse()

	fmt.Println("algo: ", *sortAlgo)

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: sort-algo [-u] [-algo=algorithm] <filename> ")
		os.Exit(1)
	}

	file, err := os.Open(flag.Args()[0])

	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}

	defer file.Close()

	uniqueLines := fileio.SortFile(file, *unique, *sortAlgo)

	for _, line := range uniqueLines {
		fmt.Println(line)
	}

}
