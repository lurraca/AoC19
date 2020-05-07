package main

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, test := range tests {
		t.Run(strconv.Itoa(test.mass), func(t *testing.T) {
			fuel := Fuel(test.mass)
			if fuel != test.fuel {
				t.Errorf("Fuel calculation is wrong for mass: %d", test.mass)
			}

		})
	}
}
