package logic

import (
	"fmt"
	"turing-machine/internal/model"
)

// Execute is single step of running code
func Execute(allStates map[int]map[string]model.Code, state int, tape *model.Tape) (*model.Tape, int, error) {
	currentState := allStates[state][tape.Symbol]

	if _, ok := allStates[currentState.NextState][tape.Symbol]; !ok && currentState.NextState != 0 {
		return nil, 0, fmt.Errorf("no code for this state: q%d %s", currentState.NextState, currentState.CurrentSymbol)
	}

	currentState = allStates[currentState.CurrentState][tape.Symbol]

	direction := currentState.Direction
	symbol := currentState.NextSymbol

	currentState = allStates[currentState.NextState][tape.Symbol]

	tape.Symbol = symbol
	switch direction {
	case "L":
		tape = tape.Last
	case "R":
		tape = tape.Next
	}

	return tape, currentState.CurrentState, nil
}
