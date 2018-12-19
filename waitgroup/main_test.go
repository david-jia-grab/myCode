package main

import "testing"

func TestSum(t *testing.T) {
	a := 1
	b := 2

	if a+b != 3 {
		t.Error("faild")
	}
}
