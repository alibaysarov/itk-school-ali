package main

import (
	"fmt"
	"math/rand"
)

func GenerateSlise(size int, max int) []int {

	length := rand.Intn(size)
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(max)
	}
	return slice
}

func GetEvenFilteredSlise(initSlise []int) []int {
	var resultSlise []int
	for _, v := range initSlise {
		if v%2 == 0 {
			resultSlise = append(resultSlise, v)
		}
	}
	return resultSlise
}
func AddElems(slise []int, elems ...int) []int {
	var newSlise []int
	newSlise = append(slise, elems...)
	return newSlise
}

func CopySlice(slise []int) []int {
	resSlise := slise[:]
	return resSlise
}

func RemoveElement(slise []int, index int) []int {
	var resSlise []int
	for i, v := range slise {
		if i != index {
			resSlise = append(resSlise, v)
		}
	}
	return resSlise
}
func main() {
	originalSlise := []int{1, 2, 3, 4}
	copiedSlise := CopySlice(originalSlise)
	originalSlise = RemoveElement(originalSlise, 3)
	fmt.Println(originalSlise)
	fmt.Println(copiedSlise)
	// fmt.Println(generateSlise(22,200))
	// fmt.Println(getEvenFilteredSlise(originalSlise))
}
