package main

import "testing"

func TestStringIntMapAddMethod(t *testing.T) {
	m := NewStringIntMap()
	initLength := 0
	m.Add("Foo Bar", 1)
	m.Add("Name", 123)
	lengthAfterInsert := len(m.data)

	if lengthAfterInsert != 2 && lengthAfterInsert <= initLength {
		t.Errorf("Test failed with params %s %s", "Foo bar", "Name")
	}
}

func TestStringIntMapRemoveMethod(t *testing.T) {
	m := NewStringIntMap()
	m.Add("Foo Bar", 1)
	m.Add("Name", 123)
	lengthAfterInsert := len(m.data)
	m.Remove("Name")
	lengthAfterRemove := len(m.data)

	if lengthAfterRemove != 1 && lengthAfterInsert <= lengthAfterRemove {
		t.Errorf("Test failed with params %s %s", "Foo bar", "Name")
	}
}

func TestStringIntMapGetMethod(t *testing.T) {
	m := NewStringIntMap()
	m.Add("Foo", 123)
	expected, expectedExists := 123, true
	result, exists := m.Get("Foo")

	if exists != expectedExists || result != expected {
		t.Errorf("Test failed with params %s", "Foo bar")
	}

}

func TestStringIntMapGetMethodWhenNotFound(t *testing.T) {
	m := NewStringIntMap()
	m.Add("Foo", 123)
	expected, expectedExists := 0, false
	result, exists := m.Get("Fo")

	if exists != expectedExists || result != expected {
		t.Errorf("Test failed with params %s", "Foo bar")
	}

}
