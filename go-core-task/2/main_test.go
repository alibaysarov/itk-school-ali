package main

import "testing"

func TestGenerateSlise(t *testing.T) {
	size := 10
	testSlise := GenerateSlise(size, 100)
	if len(testSlise) > size {
		t.Errorf("Test failed with params %d %d", size, len(testSlise))
	}
}

func TestGetEvenFilteredSlise(t *testing.T) {
	originalSlise := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := GetEvenFilteredSlise(originalSlise)
	for i, v := range expected {
		if v != result[i] {
			t.Errorf("Test failed with params %d %d %d", i, result[i], v)
		}
	}
}

func TestAddElems(t *testing.T) {
	originalSlise := []int{1, 2, 3, 4, 5, 6}
	newSlise := AddElems(originalSlise, 7, 8)
	if len(newSlise) <= len(originalSlise) {
		t.Errorf("TestAddElems failed with params: originalSlice=%v, newSlice=%v", originalSlise, newSlise)
	}
}

func TestRemoveElems(t *testing.T) {
	originalSlise := []int{1, 2, 3, 4, 5, 6}
	newSlise := RemoveElement(originalSlise, 3)
	if len(newSlise) >= len(originalSlise) {
		t.Errorf("TestAddElems failed with params: originalSlice=%v, newSlice=%v", originalSlise, newSlise)
	}
}
