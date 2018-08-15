package stack

import "testing"

func TestStack(t *testing.T) {
	var s *Stack = New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	e := s.Push("foo")
	if _, ok := e.(*StackError); !ok {
		t.Error("did not get back an error when pushing incompatible type", e)
	}
	if s.Count() != 3 {
		t.Error("pushed three, didn't find them")
	} else if v, e := s.Pop(); e != nil || v != 3 {
		t.Error("expected element not popped")
	} else if v, e := s.Peek(); e != nil || v != 2 {
		t.Error("peek broken")
	} else if s.Count() != 2 {
		t.Error("count broken.")
	}
}
