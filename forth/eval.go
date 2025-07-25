//go:build !solution

package main

import (
	"strconv"
	"strings"
)

type Evaluator struct {
	stack []int
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	return &Evaluator{}
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
		}
	}

	return e.stack, nil
}
