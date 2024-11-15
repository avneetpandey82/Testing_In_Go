package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 6
	if result != expected {
		t.Fatalf("Addition of 2+3 expected: %d but got %d", expected, result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(10, 2)
	expected := 0
	if result != expected {
		t.Fatalf("Subtract of 10-2 expected: %d but got %d", expected, result)
	}
}

func TestMul(t *testing.T) {
	result := Mul(20, 4)
	expected := 80
	if result != expected {
		t.Fatalf("Multiplication of 20*4 expected: %d but got %d", expected, result)
	}
}
