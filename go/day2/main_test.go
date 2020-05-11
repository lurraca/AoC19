package main

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{name: "2", input: []int{1, 0, 0, 0, 99}, output: []int{2, 0, 0, 0, 99}},
		{name: "6", input: []int{2, 3, 0, 3, 99}, output: []int{2, 3, 0, 6, 99}},
		{name: "9801", input: []int{2, 4, 4, 5, 99, 0}, output: []int{2, 4, 4, 5, 99, 9801}},
		{name: "30", input: []int{1, 1, 1, 4, 99, 5, 6, 0, 99}, output: []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := Intcode(test.input)
			if !reflect.DeepEqual(output, test.output) {
				t.Errorf("Intcode computating is wrong for test: %s. Expected %v, Got %v", test.name, test.output, output)
			}

		})
	}
}
