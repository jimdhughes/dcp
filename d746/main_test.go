package main

import (
	"testing"
)

func TestStackPush(t *testing.T) {
	s := ConstantStack{}
	s.push(1)
	if s.values[0] != 1  {
		t.Fail()
	}
}

func TestStackPop(t *testing.T) {
	s := ConstantStack{}
	s.push(1)
	s.push(3)
	s.push(2)
	s.push(0)
	val := s.pop()
	if val != 0 {
		t.Fail()
	}
	val = s.pop() 
	if val != 2 {
		t.Fail()
	}
}

func TestStackMax(t *testing.T) {
	s := ConstantStack{}
	s.push(1)
	s.push(5)
	s.push(3)
	s.push(4)
	max := s.max()
	if max != 5 {
		t.Fail()
	}
	s.pop()
	max = s.max()
	if max != 5 {
		t.Fail()
	}
}

func TestStackMaxWithPoppingMax(t *testing.T) {
	s := ConstantStack{}
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(3)
	s.push(5)
	max := s.max()
	if max != 5 {
		t.Fail()
	}
	s.pop()
	max = s.max()
	if max != 4 {
		t.Fail()
	}
	s.pop()
	s.pop()
	max = s.max()
	if max != 3 {
		t.Fail()
	}
}