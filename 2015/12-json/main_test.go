package main

import "testing"

func TestParseRed(t *testing.T) {
	input = []byte("[1,2,3]")
	if s := parseRed(); s != 6 {
		t.Errorf("parseRed(\"[1,2,3]\") = %d, expected 6", s)
	}

	input = []byte("[1,{\"c\":\"red\",\"b\":2},3]")
	if s := parseRed(); s != 4 {
		t.Errorf("parseRed(\"[1,{\"c\":\"red\",\"b\":2},3]\") = %d, expected 4", s)
	}

	input = []byte("{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}")
	if s := parseRed(); s != 0 {
		t.Errorf("parseRed(\"{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}\") = %d, expected 0", s)
	}

	input = []byte("[1,\"red\",5]")
	if s := parseRed(); s != 6 {
		t.Errorf("parseRed(\"[1,\"red\",5]\") = %d, expected 6", s)
	}
}
