package model

import (
	"fmt"
	"strings"
)

type State struct {
	Symbol string
	C      Code
}

type Code struct {
	CurrentState  int
	CurrentSymbol string
	NextState     int
	NextSymbol    string
	Direction     string
}

func (c *Code) String() string {
	if c == nil {
		return "<empty_code>"
	}

	return fmt.Sprintf("<cur_q:%d, cur_sym: %s, next_q:%d, next_sym: %s, dir: %s>", c.CurrentState,
		c.CurrentSymbol, c.NextState, c.NextSymbol, c.Direction)
}

/*
IsValid validate turing code:

1.Check is next state equal q1

2.Check is current state equal q1

3.Check direction (R, L, C)
*/
func (c *Code) IsValid() error {
	if c.NextState == 1 {
		return fmt.Errorf("cur_q:%d %s going to q1", c.CurrentState, c.CurrentSymbol)
	}

	if c.CurrentState == 0 {
		return fmt.Errorf("code from q0 state")
	}

	direction := strings.ToLower(c.Direction)
	if direction != "r" && direction != "c" && direction != "l" {
		return fmt.Errorf("unstandart direction: %s", c.Direction)
	}

	return nil
}
