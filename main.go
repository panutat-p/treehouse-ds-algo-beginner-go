package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 2, 3, 4}
	fmt.Println(sum(sl))
}

func sum(sl []int) int {
	if len(sl) == 1 {
		return sl[0]
	}
	return sl[0] + sum(sl[1:])

}
