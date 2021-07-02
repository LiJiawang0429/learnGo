package main

import "testing"

func TestGetHello(t *testing.T) {
	s := GetHello()
	if s != "hello" {
		t.Error("FAIL")
	}
}
