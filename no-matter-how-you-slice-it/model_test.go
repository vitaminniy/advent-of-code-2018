package main

import "testing"

func Test_newRectangle(t *testing.T) {
	const input = "#1 @ 871,327: 16x20"
	actual, err := newRectangle(input)
	if err != nil {
		t.Errorf("error occured during transforming input %s to rectangle: %v", input, err)
	}
	expected := rectangle{
		id:     1,
		x:      871,
		y:      327,
		width:  16,
		height: 20,
	}

	if actual != expected {
		t.Errorf("actual must be equal to expected:\na: %#v\ne: %#v", actual, expected)
	}
}
