package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func GetFuncType(s any) string {
	return reflect.TypeOf(s).Name()
}
func main() {
	fmt.Println(GetFuncType("13"))
	var a_10, b_10, c_10 = 10, 20, 30
	var (
		a_8 int = 1
		b_8 int = 2
		c_8 int = 3
	)

	hex1 := fmt.Sprintf("%x", a_10)
	hex2 := fmt.Sprintf("%X", b_10)
	hex3 := fmt.Sprintf("%X", c_10)

	var float_num float64 = 64.5

	var isTrue bool = true

	var z complex64 = complex(3.14, 2.0)

	var types string = fmt.Sprintf("Types %s %s %s\n %s %s %s %s\n %s %s %s\n %s %s",
		GetFuncType(a_10), GetFuncType(b_10), GetFuncType(c_10),
		GetFuncType(a_8), GetFuncType(b_8), GetFuncType(c_8),
		GetFuncType(hex1), GetFuncType(hex2), GetFuncType(hex3),
		GetFuncType(float_num), GetFuncType(isTrue), GetFuncType(z),
	)
	fmt.Println(types)

	var totalString string = fmt.Sprintf("%d %d %d %02d %02d %02d %s %s %s %g %t %v",
		a_10, b_10, c_10,
		a_8, b_8, c_8,
		hex1, hex2, hex3,
		float_num,
		isTrue,
		z,
	)
	fmt.Println(totalString)
	var runeSlice []rune = []rune(totalString)
	fmt.Println("rune ", runeSlice)

	salt := []rune("go-2024")
	mid := len(runeSlice) / 2

	var combinedRunes []rune
	combinedRunes = append(combinedRunes, runeSlice[:mid]...)
	combinedRunes = append(combinedRunes, salt...)
	combinedRunes = append(combinedRunes, runeSlice[mid:]...)
	saltyString := string(combinedRunes)
	hashedString := sha256.Sum256([]byte(saltyString))
	fmt.Printf("SHA256 hash: %x\n", hashedString)
}
