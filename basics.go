package main

import (
	"fmt"
)


func helloWolrd() {
	fmt.Println("hello world")
}

func add(x, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

func sum_char_codes(n string) {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += int(n[i])
	}

	fmt.Println(sum)
}

func main() {
	helloWolrd()
	add(15, 35)
	sum_char_codes("francielle")
}



