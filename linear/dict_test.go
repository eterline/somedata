package somedata_test

import (
	"testing"

	somedata "github.com/eterline/somedata/linear"
)

func TestDict_SetExistsDelete(t *testing.T) {
	v := somedata.NewDict[string]()

	v.Set("hello")
	if !v.Exists("hello") {
		t.Error("Exists should return true after Set")
	}

	v.Delete("hello")
	if v.Exists("hello") {
		t.Error("Exists should return false after Delete")
	}

	if v.Exists("world") {
		t.Error("Exists should return false for missing element")
	}
}

func TestDict_SizeSliceClear(t *testing.T) {
	v := somedata.NewDict[int]()
	if v.Size() != 0 {
		t.Error("Size of empty dict should be 0")
	}

	v.Set(1)
	v.Set(2)
	v.Set(3)

	if v.Size() != 3 {
		t.Errorf("Size should be 3, got %d", v.Size())
	}

	slc := v.Slice()
	found := make(map[int]bool)
	for _, x := range slc {
		found[x] = true
	}

	for _, x := range []int{1, 2, 3} {
		if !found[x] {
			t.Errorf("Slice missing element %d", x)
		}
	}

	v.Clear()
	if v.Size() != 0 {
		t.Error("Size should be 0 after Clear")
	}
	if len(v.Slice()) != 0 {
		t.Error("Slice should be empty after Clear")
	}
}

func TestDict_Intersects(t *testing.T) {
	v1 := somedata.NewDict[int]()
	v2 := somedata.NewDict[int]()

	slice, ok := v1.Intersects(v2)
	if ok {
		t.Error("Intersection of empty vocabularies should be false")
	}
	if slice != nil {
		t.Error("Intersection slice should be nil for empty vocabularies")
	}

	for i := 1; i <= 5; i++ {
		v1.Set(i)
	}
	for i := 3; i <= 7; i++ {
		v2.Set(i)
	}

	slice, ok = v1.Intersects(v2)
	if !ok {
		t.Error("Intersection should be true")
	}

	expected := map[int]bool{3: true, 4: true, 5: true}
	for _, x := range slice {
		if !expected[x] {
			t.Errorf("Unexpected intersected value: %d", x)
		}
		delete(expected, x)
	}
	if len(expected) != 0 {
		t.Errorf("Missing intersected values: %v", expected)
	}
}

func TestDict_DeleteNonExisting(t *testing.T) {
	v := somedata.NewDict[int]()
	v.Delete(999)
}

func TestDict_ExistsOnEmpty(t *testing.T) {
	v := somedata.NewDict[string]()
	if v.Exists("nothing") {
		t.Error("Exists should return false on empty dict")
	}
}
