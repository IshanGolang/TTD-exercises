package main

import (
	"reflect"
	"testing"
)

func TestEncodeDirect(t *testing.T) {
	input := []interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e'}
	expected := []struct {
		Count   int
		Element interface{}
	}{
		{4, 'a'},
		{1, 'b'},
		{2, 'c'},
		{2, 'a'},
		{1, 'd'},
		{4, 'e'},
	}

	result := encodeDirect(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: Expected %v, but got %v", expected, result)
	}
}
