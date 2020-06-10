package main

import (
	"reflect"
	"testing"
)

func TestFindInteresect(t *testing.T) {
	tests := []struct {
		name   string
		input1 Wire
		input2 Wire
		output []Point
	}{
		{"intersect on 100,0", GenerateWire("U100,R50,D50,R75,D49"), GenerateWire("R500,U40,L499,U10"), []Point{{125, 40}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			intersection := Intersection(test.input1, test.input2)
			if !reflect.DeepEqual(intersection, test.output) {
				t.Errorf("Test case: %s.Expected: %v, got: %v", test.name, test.output, intersection)
			}

		})
	}
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		name   string
		point1 Point
		point2 Point
		output int
	}{
		{"distance1", Point{20, 10}, Point{10, 10}, 10},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expectedDistance := ManhattanDistance(test.point1, test.point2)
			if !reflect.DeepEqual(expectedDistance, test.output) {
				t.Errorf("Test case: %s.Expected: %v, got: %v", test.name, test.output, expectedDistance)
			}
		})
	}
}

func TestCreateWire(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output Wire
	}{
		{"Wire 1",
			"R8,U5,L5,D3",
			Wire{
				[]Segment{
					{Point{0, 0}, Point{8, 0}},
					{Point{8, 0}, Point{8, 5}},
					{Point{8, 5}, Point{3, 5}},
					{Point{3, 5}, Point{3, 2}},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			wire := GenerateWire(test.input)
			if !reflect.DeepEqual(wire, test.output) {
				t.Errorf("Test case: %s.Expected: %v, got: %v", test.name, test.output, wire)
			}

		})
	}
}
