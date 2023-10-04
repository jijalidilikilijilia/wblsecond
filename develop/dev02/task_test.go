package main

import "testing"

func TestUnpackString(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
	}

	for _, tc := range testCases {
		result, err := UnpackString(tc.input)
		if tc.err && err == nil {
			t.Errorf("Expected error for input %q, but got none", tc.input)
		}
		if !tc.err && err != nil {
			t.Errorf("Unexpected error for input %q: %v", tc.input, err)
		}
		if result != tc.expected {
			t.Errorf("For input %q, expected %q, but got %q", tc.input, tc.expected, result)
		}
	}
}
