package main

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	encodedList := []Encoded{
		{4, 'a'},
		{1, 'b'},
		{2, 'c'},
		{2, 'a'},
		{1, 'd'},
		{4, 'e'},
	}
	expected := []interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e'}

	result := decode(encodedList)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: Expected %v, but got %v", expected, result)
	}
}
