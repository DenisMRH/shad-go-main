//go:build !solution

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Evaluator struct {
	stack   []int
	command map[string]func()
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	e := &Evaluator{
		stack:   []int{},
		command: make(map[string]func()),
	}
	addCommands(e)
	return e
}

func addCommands(e *Evaluator) {
	e.command["+"] = func() {
		fmt.Println("test")
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

		if e.command[lowerWord] != nil {
			action := e.command[lowerWord]
			action()
		}
	}

	return e.stack, nil
}
