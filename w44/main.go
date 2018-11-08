package main

import (
	"errors"
	"fmt"
)

// Stack - custom stack object
type Stack struct {
	vals []int
	maxs []int
}

// ErrStackIsEmpty - error for when a stack is empty
var ErrStackIsEmpty = errors.New("Stack is empty")

func (s *Stack) init() {
	s.vals = []int{}
	s.maxs = []int{}
}

func (s *Stack) push(val int) {
	i := len(s.vals)
	if i == 0 {
		s.maxs = append(s.maxs, val)
	} else {
		if s.maxs[i-1] > val {
			s.maxs = append(s.maxs, s.maxs[i-1])
		} else {
			s.maxs = append(s.maxs, val)
		}
	}
	s.vals = append(s.vals, val)
}

func (s *Stack) pop() (int, error) {
	i := len(s.vals) - 1
	if i == -1 {
		return 0, ErrStackIsEmpty
	}
	val := s.vals[i]
	s.vals = s.vals[:i]
	s.maxs = s.maxs[:i]
	return val, nil
}

func (s *Stack) max() (int, error) {
	i := len(s.vals) - 1
	if i == -1 {
		return 0, ErrStackIsEmpty
	}
	return s.maxs[i], nil
}

func main() {
	fmt.Println("Run `go test` to test this code out")
}
