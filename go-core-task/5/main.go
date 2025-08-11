package main

import "fmt"

func CountOccur(listArr []int, tgt int) int {
	count := 0
	for _, v := range listArr {
		if v == tgt {
			count++
		}
	}
	return count
}

func contains(slice []int, item int) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func uniqueSlice(slice []int) []int {
	set := make(map[int]struct{})
	var result []int
	for _, v := range slice {
		if _, exists := set[v]; !exists {
			set[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	slice2 := []int{3, 4, 5}

	merged := append(uniqueSlice(slice1), uniqueSlice(slice2)...)
	resMap := []int{}

	for _, v := range merged {
		if CountOccur(merged, v) > 1 && !contains(resMap, v) {
			resMap = append(resMap, v)
		}
	}
	fmt.Println(resMap)
}
