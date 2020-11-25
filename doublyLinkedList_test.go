package linkedlist

import "testing"

func TestDelete(t *testing.T) {
	l1 := New()
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)
	l1.PushBack(4)

	l2 := New()
	l2.PushBack(3)
	l2.PushBack(4)

	e := l1.GetFirst()
	l2.Delete(e)
	// l2 should not change because e is not an element of l2
	if n := l2.Length(); n != 2 {
		t.Errorf("l2.Length() = %d, want 2", n)
	}

	l1.Delete(e)
	// l1 should change because e is an element of l1
	if n := l1.Length(); n != 3 {
		t.Errorf("l2.Length() = %d, want 3", n)
	}
}

func TestNextAndBack(t *testing.T) {
	l := New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	e := l.GetFirst()
	l.Delete(e)
	if e.Value != 1 {
		t.Errorf("e.value = %d, want 1", e.Value)
	}

	if e.Next() != nil {
		t.Errorf("e.Next() != nil")
	}
	if e.Prev() != nil {
		t.Errorf("e.Prev() != nil")
	}

	l.PushBefore(8, e)
	if n := l.Length(); n != 3 {
		t.Errorf("l1.Length() = %d, want 3", n)
	}

	if l.GetFirst() == e {
		t.Errorf("l1.first() = %v, should be different", e)
	}
}


func TestPushBeforeAndAfter(t *testing.T) {
	l := New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)

	e1 := l.GetFirst().Next().Next()

	l.PushBefore(5, e1)
	val := e1.Prev().Value
	if val != 5 {
		t.Errorf("e1.Prev().Value = %d, want 5", e1.Prev().Value)
	}

	l.PushAfter(6, e1)
	val = e1.Next().Value
	if val != 6 {
		t.Errorf("e1.Prev().Value = %d, want 6", e1.Prev().Value)
	}
}

func TestSwap(t *testing.T) {
	l := New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.PushBack(5)

	e1 := l.GetFirst()
	e2 := l.GetLast()

	l.Swap(e1, e2)

	e3 := l.GetFirst()
	e4 := l.GetLast()

	if e1 != e4 {
		t.Errorf("%v should be equals to %v", e4, e1)
	}

	if e2 != e3 {
		t.Errorf("%v should be equals to %v", e3, e2)
	}

	if e3.Prev() != nil {
		t.Errorf("%v should be nil", e3)
	}

	if e4.Next() != nil {
		t.Errorf("%v should be nil", e4)
	}

	e1 = e3.Next()
	e2 = e3.Next().Next()

	l.Swap(e1, e2)

	e5 := e3.Next()
	e6 := e3.Next().Next()

	if e1 != e6 {
		t.Errorf("%v should be equals to %v", e1, e6)
	}

	if e2 != e5 {
		t.Errorf("%v should be equals to %v", e2, e5)
	}
}

func TestSwap2(t *testing.T) {
	l := New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	e1 := l.GetFirst()
	e2 := l.GetFirst().Next()

	l.Swap(e2, e1)

	e3 := l.GetFirst()
	e4 := l.GetFirst().Next()

	if e1 != e4 {
		t.Errorf("%v should be equals to %v", e4, e1)
	}

	if e2 != e3 {
		t.Errorf("%v should be equals to %v", e3, e2)
	}

	if e3.Prev() != nil {
		t.Errorf("%v should be nil", e3)
	}

	e1 = l.GetLast()
	e2 = l.GetLast().Prev()

	l.Swap(e2, e1)

	e3 = l.GetLast()
	e4 = l.GetLast().Prev()

	if e1 != e4 {
		t.Errorf("%v should be equals to %v", e1, e3)
	}

	if e2 != e3 {
		t.Errorf("%v should be equals to %v", e2, e4)
	}

	if e3.Next() != nil {
		t.Errorf("%v should be nil", e4)
	}
}