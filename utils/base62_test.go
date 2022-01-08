package utils

import "testing"

func TestNumberToBase62(t *testing.T) {
	var tests = []struct {
		input  int64
		output string
	}{
		{0, "0"},
		{1, "1"},
		{10, "A"},
		{62, "10"},
		{12345, "3D7"},
		{1234567890, "1LY7VK"},
		{123456789012345, "Z3WbxDVB"},
	}
	for _, test := range tests {
		if output := Base62(test.input); output != test.output {
			t.Errorf("%d: want %s, got %s", test.input, test.output, output)
		}
	}
}

func TestBase62ToNumber(t *testing.T) {
	var tests = []struct {
		input  string
		output int64
	}{
		{"0", 0},
		{"1", 1},
		{"A", 10},
		{"10", 62},
		{"3D7", 12345},
		{"1LY7VK", 1234567890},
		{"Z3WbxDVB", 123456789012345},
	}
	for _, test := range tests {
		if output := Base10(test.input); output != test.output {
			t.Errorf("%s: want %d, got %d", test.input, test.output, output)
		}
	}
}