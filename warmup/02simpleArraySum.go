package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	var sum int32
	for _, item := range ar {
		sum += item
	}
	return sum
}

func main() {
	sum := simpleArraySum([]int32{1, 2, 3, 4, 5})
	fmt.Println(sum)
}
