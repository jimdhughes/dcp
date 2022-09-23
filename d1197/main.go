package main

import "errors"

type Stack struct {
	values   []int
	maxValue int
}

func (s *Stack) Push(val int) {
	s.values = append(s.values, val)
	if val > s.maxValue {
		s.maxValue = val
	}
}

func (s *Stack) Pop() int {
	val := s.values[len(s.values)-1]
	s.values = s.values[0 : len(s.values)-1]
	return val
}

func (s *Stack) Max() (int, error) {
	if len(s.values) == 0 {
		return 0, errors.New("the stack is empty. Max value does not exist")
	}
	return s.maxValue, nil
}
