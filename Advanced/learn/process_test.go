package main

import "testing"

func TestCheckEven(t *testing.T) {
	i := 10
	expected := "YESs"
	res := checkEven(i)

	if expected != res {
		t.Errorf("Expected: %v, got: %v", expected, res)
	}
}
