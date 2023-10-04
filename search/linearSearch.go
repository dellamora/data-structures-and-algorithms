package main

func linear_search(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
}

func main() {
	println(linear_search([]int{1, 2, 3, 4, 5}, 8))
}
