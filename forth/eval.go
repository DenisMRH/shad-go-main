//go:build !solution

package main

import (
	"strconv"
	"strings"
)

type Evaluator struct {
	stack   []int
	command map[string]func(*Evaluator)
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	e := &Evaluator{
		stack:   []int{},
		command: make(map[string]func(*Evaluator)),
	}
	addCommands(e)
	return e
}

func addCommands(e *Evaluator) {
	e.command["+"] = func(*Evaluator) {
		e.stack[len(e.stack)-2] = e.stack[len(e.stack)-2] + e.stack[len(e.stack)-1]
		e.stack = e.stack[:len(e.stack)-1]
	}
	e.command["-"] = func(*Evaluator) {
		e.stack[len(e.stack)-2] = e.stack[len(e.stack)-2] - e.stack[len(e.stack)-1]
		e.stack = e.stack[:len(e.stack)-1]
	}
}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.
func (e *Evaluator) Process(row string) ([]int, error) {
	words := strings.Fields(row)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		if number, err := strconv.Atoi(lowerWord); err == nil {
			e.stack = append(e.stack, number)
			continue
		}

		if action, ok := e.command[lowerWord]; ok {
			action(e)
		}
	}

	return e.stack, nil
}
