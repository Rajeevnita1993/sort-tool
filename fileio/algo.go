package fileio

import (
	"crypto/rand"
	"sort"
)

// Radix Sort
func radixSort(arr []string) {
	if len(arr) <= 1 {
		return
	}

	maxLen := len(arr[0])
	for _, str := range arr {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}

	buckets := make([][]string, 256)

	for i := maxLen - 1; i >= 0; i-- {
		for _, str := range arr {
			index := 0
			if i < len(str) {
				index = int(str[i])
			}
			buckets[index] = append(buckets[index], str)
		}

		j := 0
		for k := 0; k < 256; k++ {
			for _, str := range buckets[k] {
				arr[j] = str
				j++
			}
			buckets[k] = nil
		}
	}
}

// Merge Sort
func mergeSort(arr []string, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(arr, left, mid)
		mergeSort(arr, mid+1, right)
		merge(arr, left, mid, right)
	}
}

func merge(arr []string, left, mid, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	L := make([]string, n1)
	R := make([]string, n2)

	for i := 0; i < n1; i++ {
		L[i] = arr[left+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = arr[mid+1+j]
	}

	i := 0
	j := 0
	k := left

	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

// Quick Sort
func quickSort(arr []string, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []string, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Heap Sort
func heapSort(arr []string) {
	n := len(arr)

	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// One by one extract an element from heap
	for i := n - 1; i > 0; i-- {
		// Move current root to end
		arr[0], arr[i] = arr[i], arr[0]

		// call max heapify on the reduced heap
		heapify(arr, i, 0)
	}
}

func heapify(arr []string, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func randomSort(lines []string) {
	// Get a randon hash func

	hashFunc := getRandomHashFunc()
	// Create a map to store both hash value of lines
	hashValues := make(map[string]uint64)

	// Calculate hash value of each line
	for _, line := range lines {
		hashValues[line] = hashFunc(line)
	}

	// Sort the line based on hash values
	sort.SliceStable(lines, func(i, j int) bool {
		return hashValues[lines[i]] < hashValues[lines[j]]
	})
}

func getRandomHashFunc() func(string) uint64 {
	// Generate a random seed
	var seed [8]byte
	if _, err := rand.Read(seed[:]); err != nil {
		panic(err)
	}

	// Return a hash function using random seed
	return func(s string) uint64 {
		h := uint64(5381)
		for _, c := range s {
			h = ((h << 5) + h) + uint64(c)
		}
		return h

	}
}
