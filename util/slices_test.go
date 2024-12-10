package util

import "testing"

func TestUnique(t *testing.T) {
	s := Unique([]int{1, 2, 3, 1, 2, 3})

	if len(s) != 3 || s[0] != 1 || s[1] != 2 || s[2] != 3 {
		t.Error("Expected unique array, got ", s)
	}
}
