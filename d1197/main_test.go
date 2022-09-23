package main

import "testing"

func TestPushOneValueExpectSingleValueInArray(t *testing.T) {
	s := Stack{}
	if len(s.values) > 0 {
		t.Errorf("Initialized stack should have no values")
	}
	s.Push(1)
	if len(s.values) != 1 {
		t.Errorf("More than 1 item ended up on the stack")
	}
	if s.values[0] != 1 {
		t.Errorf("Expected the value 1 to be pushed to the stack. Got %d\n", s.values[0])
	}
}

func TestPopOnlyValueExpectEmptyArray(t *testing.T) {
	s := Stack{}
	s.Push(1)
	val := s.Pop()
	if val != 1 {
		t.Errorf("Expected 1. got %d\n", val)
	}
	if len(s.values) > 0 {
		t.Errorf("Stack still has values. Expect it to be empty")
	}
}

func TestMaxValueChanges(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	if val, _ := s.Max(); val != 2 {
		t.Errorf("Expected max value to be 2. got %d\n", val)
	}
	s.Push(6)
	s.Push(3)
	if val, _ := s.Max(); val != 6 {
		t.Errorf("Expected max value to be 6. got %d\n", val)
	}
}
