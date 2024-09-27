package logic

import (
	"bufio"
	"fmt"
	"turing-machine/internal/model"
)

const (
	red   = "\033[31m"
	reset = "\033[0m"
)

/*
TapeWrite write symbol from model.Tape in bufio.Writer

CurrentCell is compare with tape to color painting in out
*/
func TapeWrite(tape, currentCell *model.Tape, writer *bufio.Writer) {
	for tape.Next != nil {
		if ColorPaint(tape, currentCell) {
			fmt.Fprintf(writer, red+"[%s]"+reset, tape.Symbol)
		} else {
			fmt.Fprintf(writer, "[%s]", tape.Symbol)
		}

		tape = tape.Next
	}
	fmt.Fprintln(writer)
}

func ColorPaint(tape, current *model.Tape) bool {
	return tape == current
}
