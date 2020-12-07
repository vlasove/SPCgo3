//main_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	var ans int
	ans = add(2, 3)
	if ans != 5 {
		t.Error("Expected 5 but has :", ans)
	}

	ans = add(2, 0)
	if ans != 2 {
		t.Error("Expected 2 but has :", ans)
	}
}

func TestSub(t *testing.T) {
	var ans int
	ans = sub(10, 2)
	if ans != 8 {
		t.Error("Expected 8 but has :", ans)
	}
}

func TestMult(t *testing.T) {
	var ans int
	ans = mult(2, 3)
	if ans != 6 {
		t.Error("Expected 6 but has :", ans)
	}
}
