package logic

import "turing-machine/internal/model"

// CreateCell fill cell with specified symbol
// insert this cell to the right and move current position of tape to the right
func CreateCell(symbol string, tape *model.Tape) *model.Tape {
	cell := model.Tape{
		Symbol: symbol,
		Last:   tape,
	}

	tape.Next = &cell
	tape = tape.Next

	return tape
}
