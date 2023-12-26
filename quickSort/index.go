package main

import "fmt"

func quickSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[len(arr)/2]
    var left, right []int

    for _, value := range arr {
        if value < pivot {
            left = append(left, value)
        } else if value > pivot {
            right = append(right, value)
        }
    }

    return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
    unsortedArray := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
    sortedArray := quickSort(unsortedArray)
    fmt.Println(sortedArray)
}
