package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	var sum int64

	for _, item := range ar {
		sum += item
	}
	return sum
}

func main() {
	sum := aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
	fmt.Println(sum)
}
