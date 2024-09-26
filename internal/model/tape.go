package model

import "fmt"

type Tape struct {
	Symbol string
	Last   *Tape
	Next   *Tape
}

func NewTape(symbol string) *Tape {
	return &Tape{
		Symbol: symbol,
		Last:   nil,
		Next:   nil,
	}
}

func (t *Tape) String() string {
	if t == nil {
		return "<empty_tape>"
	}

	return fmt.Sprintf("<sym: %s, last_addr: %p, next_addr: %p>", t.Symbol, t.Last, t.Next)
}
