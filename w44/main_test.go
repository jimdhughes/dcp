package main

import (
	"os"
	"testing"
)

var s Stack

func TestMain(m *testing.M) {
	s := &Stack{}
	s.init()
	os.Exit(m.Run())
}

func TestPopEmptyStack(t *testing.T) {
	_, err := s.pop()
	s.init()
	if err != ErrStackIsEmpty {
		t.Errorf("Expected ErrStackIsEmpty")
	}
}

func TestMaxEmptyStack(t *testing.T) {
	_, err := s.max()
	if err != ErrStackIsEmpty {
		t.Errorf("Expected ErrStackIsEmpty")
	}
}

func TestPushStack(t *testing.T) {
	s.init()
	s.push(1)
	val, err := s.pop()
	if err != nil {
		t.Fail()
	}
	if val != 1 {
		t.Fail()
	}
}

func TestMax(t *testing.T) {
	var max int
	var err error
	s.init()
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(5)

	max, err = s.max()
	if err != nil {
		t.Error(err)
	}
	if max != 5 {
		t.Errorf("Expected max to be %d got %d", 5, max)
	}
	s.pop()
	max, err = s.max()
	if err != nil {
		t.Error(err)
	}
	if max != 3 {
		t.Errorf("Expecte max to be %d got %d", 3, max)
	}

}
