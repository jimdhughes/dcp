package main

import (
	"errors"
	"fmt"
	"strconv"
)

type ReversePolishNotation struct {
	Input []string
	stack []int
}

func getDigit(s string) (int, error) {
	return strconv.Atoi(s)
}

func (r *ReversePolishNotation) evaluateOperand(op string) error {
	if len(r.stack) < 2 {
		return errors.New("Not enough items to evaluate")
	}
	values := r.stack[len(r.stack)-2:]
	r.stack = r.stack[0:len(r.stack)-2]
	switch op {
		case "+":
			r.stack = append(r.stack, values[0] + values[1])
			return nil
		case "-":
				r.stack = append(r.stack, values[0] - values[1])
				return nil
		case "*":
			r.stack = append(r.stack, values[0] * values[1])
				return nil
		case "/":
			r.stack = append(r.stack, values[0] / values[1])
				return nil
		default:
			return errors.New(fmt.Sprintf("Invalid operand: %s", op))
	}
}

func (r *ReversePolishNotation) handleEvaluation() (int, error){
	for _, s := range r.Input {
		intValue, err := getDigit(s)
		if err != nil {
			switch s {
			case "+", "-", "/", "*":
				err := r.evaluateOperand(s)
				if err != nil {
					return 0, errors.New("Invalid input. Unable to evaluate")
				}
				break;
			default :
				return 0, errors.New("Unrecognized character "+s)
			}
		}
		if err == nil {
			r.stack = append(r.stack, intValue)
		}
	}
	
	if len(r.stack) > 1 {
		return 0, errors.New("Invalid input.")
	}
	return r.stack[0], nil
}

func (r *ReversePolishNotation) Calculate(s []string) (int, error) {
	r.Input = s
	return r.handleEvaluation()
}