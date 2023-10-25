package main

import (
	"reflect"
	"testing"
)

func TestEncodeModified(t *testing.T) {
	input := []interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e'}
	expected := []interface{}{
		struct {
			Count   int
			Element interface{}
		}{
			Count:   4,
			Element: 'a',
		},
		'b',
		struct {
			Count   int
			Element interface{}
		}{
			Count:   2,
			Element: 'c',
		},
		struct {
			Count   int
			Element interface{}
		}{
			Count:   2,
			Element: 'a',
		},
		'd',
		struct {
			Count   int
			Element interface{}
		}{
			Count:   4,
			Element: 'e',
		},
	}

	result := encodeModified(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Test failed: Expected %v, but got %v", expected, result)
	}
}
