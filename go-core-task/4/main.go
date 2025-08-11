package main

import "fmt"

func CountOccur(listArr []string, tgt string) int {
	count := 0
	for _, v := range listArr {
		if v == tgt {
			count++
		}
	}
	return count
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func uniqueSlice(slice []string) []string {
	set := make(map[string]struct{})
	var result []string
	for _, v := range slice {
		if _, exists := set[v]; !exists {
			set[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	merged := append(uniqueSlice(slice1), uniqueSlice(slice2)...)
	resMap := []string{}

	for _, v := range merged {
		if CountOccur(merged, v) > 0 && CountOccur(merged, v) < 2 && !contains(resMap, v) {
			resMap = append(resMap, v)
		}
	}
	fmt.Println(resMap)
}
