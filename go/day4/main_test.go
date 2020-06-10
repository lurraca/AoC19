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
		{name: "is 111111 candidate", input: "111111", output: true},
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
		{"111111", "111111", true},
		{"123456", "123456", false},
		{"123455", "123455", true},
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
