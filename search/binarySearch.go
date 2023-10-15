package main

import "fmt"

func binarySearch(array []int, targ int) bool {
	lo := 0
	hi := len(array)

	for lo < hi {
		mid := lo + (hi-lo)/2
		value := array[mid]

		if value == targ {
			return true
		} else if value < targ {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return false
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 3))
}
