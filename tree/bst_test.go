package somedata_test

import (
	"reflect"
	"strconv"
	"testing"

	somedata "github.com/eterline/somedata/tree"
)

func TestThreeBST_InsertAndInOrder(t *testing.T) {
	tree := somedata.NewThreeBST[int, any]()

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v, strconv.Itoa(v))
	}

	if tree.Size() != len(values) {
		t.Fatalf("expected size %d, got %d", len(values), tree.Size())
	}

	got := []int{}
	tree.InOrder(func(v int) { got = append(got, v) })

	expected := []int{50, 30, 20, 40, 70, 60, 80}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("inOrder: expected %v, got %v", expected, got)
	}
}

func TestThreeBST_MinMax(t *testing.T) {
	tree := somedata.NewThreeBST[int, any]()

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v, strconv.Itoa(v))
	}

	min, ok := tree.Min()
	if !ok || min != 20 {
		t.Fatalf("expected min 20, got %v", min)
	}

	max, ok := tree.Max()
	if !ok || max != 80 {
		t.Fatalf("expected max 80, got %v", max)
	}
}

func TestThreeBST_Delete(t *testing.T) {
	tree := somedata.NewThreeBST[int, any]()

	values := []int{50, 30, 70, 20, 40, 60, 80}
	for _, v := range values {
		tree.Insert(v, strconv.Itoa(v))
	}

	deletes := []int{20, 70, 50}
	for _, v := range deletes {
		ok := tree.Delete(v)
		if !ok {
			t.Fatalf("expected Delete(%d) = true, got false", v)
		}
	}

	expectedSize := len(values) - len(deletes)
	if tree.Size() != expectedSize {
		t.Fatalf("expected size %d, got %d", expectedSize, tree.Size())
	}

	got := []int{}
	tree.InOrder(func(v int) { got = append(got, v) })
	expected := []int{30, 40, 60, 80}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("inOrder after delete: expected %v, got %v", expected, got)
	}
}

func TestThreeBST_DeleteNonExistent(t *testing.T) {
	tree := somedata.NewThreeBST[int, any]()
	tree.Insert(10, "hello")
	tree.Insert(20, "there")

	ok := tree.Delete(99)
	if ok {
		t.Fatalf("expected Delete(99) = false, got true")
	}

	if tree.Size() != 2 {
		t.Fatalf("expected size 2, got %d", tree.Size())
	}
}
