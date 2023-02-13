package main

import (
	"testing"
)

func TestContains(t *testing.T) {
	l := []interface{}{"a", 'b', 1, 2.0, true}
	for _, e := range l {
		if !contains(l, e) {
			t.Errorf("contains(%v, %v) = false, want true", l, e)
		}
	}
}
