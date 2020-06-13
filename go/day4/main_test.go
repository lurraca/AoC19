package main

import (
	"reflect"
	"testing"
)

func TestIsPasswordCandidate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		//{name: "is 111111 candidate", input: "111111", output: true}, IS NOT TRUE FOR PART 2
		{name: "is 123123 candidate", input: "123123", output: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isPasswordCandidate := isPasswordCandidate(test.input)
			if !reflect.DeepEqual(isPasswordCandidate, test.output) {
				t.Errorf("Test case: %s.Expected: %v, got: %v", test.name, test.output, isPasswordCandidate)
			}
		})
	}
}

func TestIsAdjacentMatch(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{"223334", "223334", true},
		{"144559", "144559", true},
		{"155678", "155678", true},
		{"122444", "122444", true},
		{"111111", "111111", false},
		{"123456", "123456", false},
		{"123455", "123455", true},
		{"112233", "112233", true},
		{"123444", "123444", false},
		{"111122", "111122", true},
		{"144445", "144445", false},
		{"144455", "144455", true},
		{"144566", "144566", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isAdjacentMatch := isAdjacentMatch(test.input)
			if !reflect.DeepEqual(isAdjacentMatch, test.output) {
				t.Errorf("Test case: %s.Expected: %v, got: %v", test.name, test.output, isAdjacentMatch)
			}
		})
	}
}

func TestDecreases(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{"111111", "111111", false},
		{"123456", "123456", false},
		{"123455", "123455", false},
		{"123655", "123655", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			decreases := decreases(test.input)
			if !reflect.DeepEqual(decreases, test.output) {
				t.Errorf("Test case: %s.Expected: %v, got: %v", test.name, test.output, decreases)
			}
		})
	}
}
