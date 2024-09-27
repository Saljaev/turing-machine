package logic

import "turing-machine/internal/model"

// SearchStart return the leftmost cell of tape
func SearchStart(tape *model.Tape) *model.Tape {
	if tape.Last != nil {
		tape = tape.Last
		SearchStart(tape)

	}
	return tape
}
