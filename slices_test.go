package main

import "testing"

func TestCompareIntSlices(t *testing.T) {
	a := []int{1, 2}
	b := []int{0}
	if compareIntSlices(a, b) != false {
		t.Errorf("compareMemories(a, b) = %v; want %v", compareIntSlices(a, b), false)
	}

	b = []int{0, 2}
	if compareIntSlices(a, b) != false {
		t.Errorf("compareMemories(a, b) = %v; want %v", compareIntSlices(a, b), false)
	}

	b = []int{1, 2}
	if compareIntSlices(a, b) != true {
		t.Errorf("compareMemories(a, b) = %v; want %v", compareIntSlices(a, b), true)
	}
}
