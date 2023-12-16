package main

import "testing"

func TestHashing(t *testing.T) {
	h := hash("HASH")
	if h != 52 {
		t.Errorf("Expected 52, got %d", h)
	}

	h = hash("rn")
	if h != 0 {
		t.Errorf("Expected 0, got %d", h)
	}
}

func TestRemoveLens(t *testing.T) {
	boxes := boxList{
		1: []lense{
			{"C", 1},
		},
		8: []lense{
			{"A", 1},
			{"B", 2},
		},
	}

	boxes = removeLens(boxes, "A")
	if len(boxes[8]) != 1 {
		t.Errorf("Expected 1, got %d", len(boxes[8]))
	}

	boxes = removeLens(boxes, "B")
	if _, ok := boxes[8]; ok {
		t.Error("Empty box should be removed but still exists")
	}

	if len(boxes[1]) != 1 {
		t.Error("Box 1 should not be affected")
	}
}
