package main

import "testing"

func TestNextPassword(t *testing.T) {
	start := []byte("xyy")
	if incrementPassword(&start); string(start) != "xyz" {
		t.Fatalf("expected password xyz, got %s", string(start))
	}
	if incrementPassword(&start); string(start) != "xza" {
		t.Fatalf("expected password xza, got %s", string(start))
	}
	if incrementPassword(&start); string(start) != "xzb" {
		t.Fatalf("expected password xzb, got %s", string(start))
	}
}

func TestIsValidPassword(t *testing.T) {
	if isValidPassword([]byte("hijklmmn")) {
		t.Error("hijklmmn should not be valid")
	}
	if isValidPassword([]byte("abbceffg")) {
		t.Error("abbceffg should not be valid")
	}
	if isValidPassword([]byte("abbcegjk")) {
		t.Error("abbcegjk should not be valid")
	}

	if !isValidPassword([]byte("abcdffaa")) {
		t.Error("abcdffaa should be valid")
	}

	if !isValidPassword([]byte("ghjaabcc")) {
		t.Error("ghjaabcc should be valid")
	}
}
