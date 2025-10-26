package somedata_test

import (
	"testing"

	"github.com/eterline/somedata"
)

func newTestMatrix2x2() (somedata.Matrix[int], somedata.Matrix[int]) {
	a := somedata.NewMatrix2[int](2, 2)
	b := somedata.NewMatrix2[int](2, 2)

	// Matix A
	a.Set(1, 0, 0)
	a.Set(2, 0, 1)
	a.Set(3, 1, 0)
	a.Set(4, 1, 1)

	// Matix B
	b.Set(5, 0, 0)
	b.Set(6, 0, 1)
	b.Set(7, 1, 0)
	b.Set(8, 1, 1)

	return a, b
}

func TestShapeEquals(t *testing.T) {
	a, b := newTestMatrix2x2()
	if !a.ShapeEquals(b) {
		t.Fatalf("expected matrices to have same shape")
	}
}

func TestAdd(t *testing.T) {
	a, b := newTestMatrix2x2()

	sum, err := a.Add(b)
	if err != nil {
		t.Fatalf("Add() returned error: %v", err)
	}

	expected := []int{6, 8, 10, 12}
	got := sum.Flatten()
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("Add mismatch at %d: got %v, want %v", i, got[i], v)
		}
	}
}

func TestSub(t *testing.T) {
	a, b := newTestMatrix2x2()

	diff, err := a.Sub(b)
	if err != nil {
		t.Fatalf("Sub() returned error: %v", err)
	}

	expected := []int{-4, -4, -4, -4}
	got := diff.Flatten()
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("Sub mismatch at %d: got %v, want %v", i, got[i], v)
		}
	}
}

func TestMulHadamard(t *testing.T) {
	a, b := newTestMatrix2x2()

	prod, err := a.MulHadamard(b)
	if err != nil {
		t.Fatalf("MulHadamard() returned error: %v", err)
	}

	expected := []int{5, 12, 21, 32}
	got := prod.Flatten()
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("Hadamard mismatch at %d: got %v, want %v", i, got[i], v)
		}
	}
}

func TestScale(t *testing.T) {
	a, _ := newTestMatrix2x2()

	scaled := a.Scale(2)
	expected := []int{2, 4, 6, 8}
	got := scaled.Flatten()
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("Scale mismatch at %d: got %v, want %v", i, got[i], v)
		}
	}
}

func TestEquals(t *testing.T) {
	a, b := newTestMatrix2x2()

	if a.Equals(b) {
		t.Errorf("Matrices should not be equal (different values)")
	}

	c, _ := a.Add(somedata.NewMatrix2[int](2, 2))
	if !a.Equals(c) {
		t.Errorf("Expected equal matrices, got not equal")
	}
}

func TestZero(t *testing.T) {
	a, _ := newTestMatrix2x2()
	a.Zero()

	got := a.Flatten()
	for i, v := range got {
		if v != 0 {
			t.Errorf("Zero failed at index %d: got %v", i, v)
		}
	}
}

func TestShape(t *testing.T) {
	a, _ := newTestMatrix2x2()
	shape := a.Shape()

	if len(shape) != 2 || shape[0] != 2 || shape[1] != 2 {
		t.Errorf("expected shape [2,2], got %v", shape)
	}
}

func TestRank(t *testing.T) {
	a, _ := newTestMatrix2x2()
	if r := a.Rank(); r != 2 {
		t.Errorf("expected rank 2, got %v", r)
	}
}

func TestSize(t *testing.T) {
	a, _ := newTestMatrix2x2()
	if s := a.Size(); s != 4 {
		t.Errorf("expected size 4, got %v", s)
	}
}
