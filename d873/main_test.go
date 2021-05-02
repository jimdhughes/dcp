package main

import (
	"testing"
)

func TestEvaluateOperand_ValidInput(t *testing.T) {
	r := ReversePolishNotation{}
	r.stack=[]int{1,2}
	err := r.evaluateOperand("+")
	if err != nil {
		t.Error(err)
	}
	if r.stack[0] != 3 {
		t.Errorf("Expected 3, got %d\n", r.stack[0])
	}
}

func TestEvaluateOperand_InvalidInput(t * testing.T) {
	r := ReversePolishNotation{}
	r.stack=[]int{1,2}
	err := r.evaluateOperand("x")
	if err == nil {
		t.Errorf("Expected error due to invalid operand. Failing test")
	}
}

func TestEvaluateOperand_NotEnoughInputs (t *testing.T) {
	r := ReversePolishNotation{}
	r.stack=[]int{1}
	err := r.evaluateOperand("+")
	if err == nil {
		t.Errorf("Expected error due to invalid stack length. Failing test")
	}
}

func TestHandleEvaluation_ValidInput(t *testing.T) {
	r := ReversePolishNotation{}
	r.Input = []string{"5", "3", "+"}
	val, err := r.handleEvaluation()
	if err != nil {
		t.Fail()
	}
	if val != 8 {
		t.Errorf("Expected 8, got: %d", val)
	}
}

func TestHandleEvaluation_ValidComplexInput(t *testing.T) {
	r := ReversePolishNotation{}
	r.Input = []string{"15", "7", "1", "1", "+", "-", "/", "3", "*", "2", "1", "1", "+", "+", "-"}
	val, err := r.handleEvaluation()
	if err != nil {
		t.Fail()
	}
	if val != 5 {
		t.Errorf("Expected 5, got: %d", val)
	}
}