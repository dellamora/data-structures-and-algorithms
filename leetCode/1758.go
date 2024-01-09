package main


func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}	

func minOperations(s string) int {
	var count int
	for i, v := range s {
		if i % 2 == 0 && v == '1' {
			count++
		} else if i % 2 == 1 && v == '0' {
			count++
		}
	}
	return min(count, len(s) - count)
}

