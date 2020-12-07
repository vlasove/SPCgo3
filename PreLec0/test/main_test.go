package main

import "testing"

func TestAdd(t *testing.T) {
	var ans int
	ans = add(2, 3)
	if ans != 5 {
		t.Error("Expected 5 has : ", ans)
	}
}
