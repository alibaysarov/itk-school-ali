package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() StringIntMap {
	return StringIntMap{
		data: make(map[string]int),
	}
}

func (testMap *StringIntMap) Add(key string, value int) {
	testMap.data[key] = value
}

func (testMap *StringIntMap) Remove(key string) {
	delete(testMap.data, key)
}

func (testMap *StringIntMap) Exists(key string) bool {
	_, exists := testMap.data[key]
	return exists
}

func (testMap *StringIntMap) Get(key string) (int, bool) {
	val, exists := testMap.data[key]
	return val, exists
}

func (testMap *StringIntMap) ToString() string {
	str := ""
	for key, value := range testMap.data {
		str += fmt.Sprintf("(%s %d) ", key, value)
	}
	return str
}

func (testMap *StringIntMap) Copy() map[string]int {
	copy := make(map[string]int, len(testMap.data))

	for key, value := range testMap.data {
		copy[key] = value
	}
	return copy
}
