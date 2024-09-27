package logic

import "turing-machine/internal/model"

// GenerateEdge extends tape in left/right side by add zero cell
func GenerateEdge(tape *model.Tape, zeroCell string) *model.Tape {
	if tape.Last == nil {
		cell := model.Tape{
			Symbol: zeroCell,
			Next:   tape,
		}

		tape.Last = &cell
	}
	if tape.Next == nil {
		cell := model.Tape{
			Symbol: zeroCell,
			Last:   tape,
		}

		tape.Next = &cell
	}

	return tape
}
