package somedata_test

import (
	"testing"

	somedata "github.com/eterline/somedata/linear"
)

func TestLinkedList_NewLinkedList(t *testing.T) {
	ll := somedata.NewLinkedList[int]()
	if ll == nil {
		t.Fatal("NewLinkedList returned nil")
	}
	if ll.Size() != 0 {
		t.Errorf("NewLinkedList size should be 0, got %d", ll.Size())
	}
}

func TestLinkedList_AddHeadTail_Size(t *testing.T) {
	ll := somedata.NewLinkedList[string]()

	// Добавление элементов
	ll.Add("first")
	ll.Add("second")
	ll.Add("third")

	if ll.Size() != 3 {
		t.Errorf("Size expected 3, got %d", ll.Size())
	}

	head, ok := ll.Head()
	if !ok || head != "third" {
		t.Errorf("Head expected 'third', got '%v'", head)
	}

	tail, ok := ll.Tail()
	if !ok || tail != "first" {
		t.Errorf("Tail expected 'first', got '%v'", tail)
	}
}

func TestLinkedList_Pop(t *testing.T) {
	ll := somedata.NewLinkedList[int]()

	// Pop пустого списка
	_, ok := ll.Pop()
	if ok {
		t.Error("Pop from empty list should return false")
	}

	// Добавление и Pop
	ll.Add(10)
	ll.Add(20)
	ll.Add(30)

	val, ok := ll.Pop()
	if !ok || val != 30 {
		t.Errorf("Expected 30 from Pop, got %v", val)
	}

	if ll.Size() != 2 {
		t.Errorf("Size expected 2 after Pop, got %d", ll.Size())
	}

	val, ok = ll.Pop()
	if !ok || val != 20 {
		t.Errorf("Expected 20 from Pop, got %v", val)
	}

	val, ok = ll.Pop()
	if !ok || val != 10 {
		t.Errorf("Expected 10 from Pop, got %v", val)
	}

	// Список пуст
	_, ok = ll.Pop()
	if ok {
		t.Error("Pop from empty list should return false")
	}
}

func TestLinkedList_Slice(t *testing.T) {
	ll := somedata.NewLinkedList[int]()
	if sl := ll.Slice(); sl != nil {
		t.Error("Slice on empty list should return nil")
	}

	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	slice := ll.Slice()
	expected := []int{1, 2, 3} // добавление пушит в голову, поэтому Slice возвращает обратный порядок

	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Slice[%d] expected %d, got %d", i, v, slice[i])
		}
	}
}

func TestLinkedList_Clear(t *testing.T) {
	ll := somedata.NewLinkedList[string]()
	ll.Add("a")
	ll.Add("b")
	ll.Add("c")

	ll.Clear()
	if ll.Size() != 0 {
		t.Errorf("Size after Clear expected 0, got %d", ll.Size())
	}

	head, ok := ll.Head()
	if ok {
		t.Errorf("Head after Clear should return false, got %v", head)
	}

	tail, ok := ll.Tail()
	if ok {
		t.Errorf("Tail after Clear should return false, got %v", tail)
	}

	if sl := ll.Slice(); sl != nil {
		t.Error("Slice after Clear should return nil")
	}
}

func TestLinkedList_NilPointerPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic on nil linked list pointer")
		}
	}()

	var ll *somedata.LinkedList[int]
	ll.Add(1) // должно паниковать
}
