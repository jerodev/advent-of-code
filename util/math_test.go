package util

import "testing"

func TestNthDigit(t *testing.T) {
	d := NthDigit(1234, 1)
	if d != 4 {
		t.Errorf("1st digit of 1234 is 4, got %d", d)
	}

	d = NthDigit(1234, 2)
	if d != 3 {
		t.Errorf("2nd digit of 1234 is 3, got %d", d)
	}

	d = NthDigit(1234, 3)
	if d != 2 {
		t.Errorf("3rd digit of 1234 is 2, got %d", d)
	}

	d = NthDigit(1234, 4)
	if d != 1 {
		t.Errorf("4th digit of 1234 is 1, got %d", d)
	}

	d = NthDigit(1234, 5)
	if d != 0 {
		t.Errorf("5th digit of 1234 is 0, got %d", d)
	}
}
