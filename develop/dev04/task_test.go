package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		words    []string
		expected map[string][]string
	}{
		{
			words: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			words:    []string{"hello", "world"},
			expected: map[string][]string{},
		},
		{
			words:    []string{},
			expected: map[string][]string{},
		},
	}

	for _, test := range tests {
		result := FindAnagrams(&test.words)
		if !reflect.DeepEqual(*result, test.expected) {
			t.Errorf("For words %v, expected %v, but got %v", test.words, test.expected, *result)
		}
	}
}
