package somedata_test

import (
	"sync"
	"testing"

	somedata "github.com/eterline/somedata/ring_buffer"
)

func TestRingBuffer_AddAndGet(t *testing.T) {
	rb := somedata.NewRingBuffer[int](5)

	for i := 1; i <= 3; i++ {
		rb.Add(i)
	}

	if rb.Len() != 3 {
		t.Errorf("expected len 3, got %d", rb.Len())
	}

	for i := 0; i < 3; i++ {
		v, ok := rb.Get(i)
		if !ok {
			t.Errorf("expected element at index %d to exist", i)
		}
		if v != i+1 {
			t.Errorf("expected %d, got %d", i+1, v)
		}
	}

	for i := 4; i <= 7; i++ {
		rb.Add(i)
	}

	if rb.Len() != 5 {
		t.Errorf("expected len 5, got %d", rb.Len())
	}

	expected := []int{3, 4, 5, 6, 7}
	values := rb.Slice()
	for i, v := range expected {
		if values[i] != v {
			t.Errorf("expected values[%d] = %d, got %d", i, v, values[i])
		}
	}
}

func TestRingBuffer_Contains(t *testing.T) {
	rb := somedata.NewRingBuffer[string](3)
	rb.Add("a")
	rb.Add("b")
	rb.Add("c")

	if !rb.Contains("b") {
		t.Errorf("expected to contain 'b'")
	}
	if rb.Contains("d") {
		t.Errorf("did not expect to contain 'd'")
	}

	rb.Add("d") // перезаписываем "a"
	if rb.Contains("a") {
		t.Errorf("did not expect to contain 'a' after overwrite")
	}
}

func TestRingBuffer_Clear(t *testing.T) {
	rb := somedata.NewRingBuffer[int](3)
	for i := 1; i <= 4; i++ {
		rb.Add(i)
	}

	rb.Clear()
	if rb.Len() != 0 {
		t.Errorf("expected len 0 after clear, got %d", rb.Len())
	}

	for i := 0; i < 4; i++ {
		_, ok := rb.Get(i)
		if ok {
			t.Errorf("expected no element after clear at index %d", i)
		}
	}
}

func TestRingBuffer_FullAndCap(t *testing.T) {
	rb := somedata.NewRingBuffer[int](3)
	if rb.Full() {
		t.Errorf("expected buffer not full initially")
	}

	rb.Add(1)
	rb.Add(2)
	if rb.Full() {
		t.Errorf("expected buffer not full yet")
	}

	rb.Add(3)
	if !rb.Full() {
		t.Errorf("expected buffer full now")
	}

	if rb.Cap() != 3 {
		t.Errorf("expected cap 3, got %d", rb.Cap())
	}
}

func TestSyncRingBuffer_ConcurrentAccess(t *testing.T) {
	rb := somedata.NewRingBuffer[int](10)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			rb.Add(v)
		}(i)
	}

	wg.Wait()

	if rb.Len() != 10 {
		t.Errorf("expected len 10 after concurrent writes, got %d", rb.Len())
	}

	slice := rb.Slice()
	for _, v := range slice {
		if v < 90 || v > 99 {
			t.Errorf("expected element in 90..99, got %d", v)
		}
	}
}

func TestRingBuffer_Boundaries(t *testing.T) {
	rb := somedata.NewRingBuffer[int](2)
	_, ok := rb.Get(-1)
	if ok {
		t.Errorf("expected Get(-1) to fail")
	}

	_, ok = rb.Get(2)
	if ok {
		t.Errorf("expected Get(2) to fail")
	}

	rb.Add(1)
	rb.Add(2)
	rb.Add(3)

	v, ok := rb.Get(0)
	if !ok || v != 2 {
		t.Errorf("expected first element = 2 after overwrite, got %d", v)
	}
}
