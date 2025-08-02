//go:build !solution

package main

import (
	"errors"
	"strconv"
	"strings"
)

var underCommand = make(map[string][]string)

type Evaluator struct {
	stack   []int
	command map[string]func(*Evaluator) error
	underCommand map[string][]string
}

// NewEvaluator creates evaluator.
func NewEvaluator() *Evaluator {
	e := &Evaluator{
		stack:   []int{},
		command: make(map[string]func(*Evaluator) error),
	}
	addCommands(e)
	return e
}

func addCommands(e *Evaluator) {
	e.command["+"] = func(*Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("для выполнения этой команды требуется чтобы в стеке было более 2 чисел")
		}
		e.stack[len(e.stack)-2] = e.stack[len(e.stack)-2] + e.stack[len(e.stack)-1]
		e.stack = e.stack[:len(e.stack)-1]
		return nil
	}
	e.command["-"] = func(*Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("для выполнения этой команды требуется чтобы в стеке было более 2 чисел")
		}
		e.stack[len(e.stack)-2] = e.stack[len(e.stack)-2] - e.stack[len(e.stack)-1]
		e.stack = e.stack[:len(e.stack)-1]
		return nil
	}
	e.command["*"] = func(*Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("для выполнения этой команды требуется чтобы в стеке было более 2 чисел")
		}
		e.stack[len(e.stack)-2] = e.stack[len(e.stack)-2] * e.stack[len(e.stack)-1]
		e.stack = e.stack[:len(e.stack)-1]
		return nil
	}
	e.command["/"] = func(*Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("для выполнения этой команды требуется чтобы в стеке было более 2 чисел")
		}
		if e.stack[len(e.stack)-1] != 0 {
			e.stack[len(e.stack)-2] = e.stack[len(e.stack)-2] / e.stack[len(e.stack)-1]
			e.stack = e.stack[:len(e.stack)-1]
			return nil
		} else {
			return errors.New("на ноль делить нельзя")
		}
	}

	e.command["swap"] = func(*Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("для выполнения этой команды требуется чтобы в стеке было хотябы два числа")
		}
		aftLst := e.stack[len(e.stack)-2]
		e.stack[len(e.stack)-2] = e.stack[len(e.stack)-1]
		e.stack[len(e.stack)-1] = aftLst
		return nil
	}

	e.command["over"] = func(*Evaluator) error {
		if len(e.stack) < 2 {
			return errors.New("для выполнения этой команды требуется чтобы в стеке было хотябы одно число")
		}
		e.stack = append(e.stack, e.stack[len(e.stack)-2])
		return nil
	}

	e.command["drop"] = func(*Evaluator) error {
		if len(e.stack) < 1 {
			return errors.New("для выполнения этой команды требуется чтобы в стеке было хотябы одно число")
		}
		e.stack = e.stack[:len(e.stack)-1]
		return nil
	}
	e.command["dup"] = func(*Evaluator) error {
		if len(e.stack) < 1 {
			return errors.New("для выполнения этой команды требуется чтобы в стеке было хотябы одно число")
		}
		e.stack = append(e.stack, e.stack[len(e.stack)-1])
		return nil
	}

}

// Process evaluates sequence of words or definition.
//
// Returns resulting stack state and an error.
func (e *Evaluator) Process(row string) ([]int, error) {
	words := strings.Fields(row)
	if words[0] == ":" {
		keyWord := strings.ToLower(words[1])
		if _, err := strconv.Atoi(keyWord); err != nil {
			underCommand[keyWord] = declarationCommand(words)
		} else {
			return e.stack, errors.New("нельзя переопределять числа")
		}
	} else {
	firstfor:
		for _, word := range words {

			lowerWord := strings.ToLower(word)

			if number, err := strconv.Atoi(lowerWord); err == nil {
				e.stack = append(e.stack, number)
				continue
			}

			if underCommand, ok := underCommand[lowerWord]; ok {
				for _, word := range underCommand {
					lowerWord := strings.ToLower(word)
					if number, err := strconv.Atoi(lowerWord); err == nil {
						e.stack = append(e.stack, number)
						continue
					}
					if action, ok := e.command[lowerWord]; ok {
						if err := action(e); err != nil {
							return e.stack, err
						}
					}

				}
				continue firstfor
			}

			if action, ok2 := e.command[lowerWord]; ok2 {
				if err := action(e); err != nil {
					return e.stack, err
				}
			}
			_, ok := e.command[lowerWord]
			_, ok1 := underCommand[lowerWord]
			if !ok && !ok1 {
				return e.stack, errors.New("нет такой комманды: " + lowerWord)
			}

		}
	}

	return e.stack, nil
}

func declarationCommand(words []string) (rowUnderCommand []string) {
	for i := 2; i < len(words); i++ {
		lowerWord := strings.ToLower(words[i])
		if _, ok := underCommand[lowerWord]; !ok {
			rowUnderCommand = append(rowUnderCommand, lowerWord)
		} else {
			rowUnderCommand = append(rowUnderCommand, underCommand[lowerWord]...)
		}
	}
	return
}
