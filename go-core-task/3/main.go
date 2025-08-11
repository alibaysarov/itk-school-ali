package main

import "fmt"

func main() {
	m := NewStringIntMap()
	m.Add("ali", 23)
	m.Add("Dev", 25)
	fmt.Println(m.ToString())
}
