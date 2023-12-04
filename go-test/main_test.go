package main

import "testing"

var tests = []struct {
	name string
	dividend float32
	divisor float32
	expected float32
	isErr bool
}{ // test name, x, y, expected result, error present?
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-five", 100.0, 20.0, 5.0, false},
}

// all tests start with T otherwise ignored
func TestDivisionGeneral(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected error did not occur")
			}
		} else {
			if err != nil {
				t.Error("Got an unexpected error", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, got)
		}
	}
}

// 2 tests  below are the same as 1 test above
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 5.0)
	if err != nil {
		t.Error("Unexpected error")
	}
}

func TestZeroDivide(t *testing.T) {
	_, err := divide(10.0, 0)
	if err == nil {
		t.Error("Did not get an error on zero division")
	}
}