package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	var aR, bR int32

	for ind, _ := range a {
		if a[ind] > b[ind] {
			aR++
		} else if a[ind] < b[ind] {
			bR++
		}
	}

	return []int32{aR, bR}

}

func main() {
	a := []int32{1, 2, 3, 4, 5}
	b := []int32{1, 2, 5, 1, 3}
	res := compareTriplets(a, b)
	fmt.Println(res)
}
