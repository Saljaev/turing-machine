package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"turing-machine/internal/model"
)

const (
	StandartLength = 6
	FullCell       = "1"
	ZeroCell       = "0"
	Red            = "\033[31m"
	Reset          = "\033[0m"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	f, err := os.ReadFile("code.txt")
	if err != nil {
		log.Fatalf("Failed to open file with Turing code: %v", err)
	}

	// All code from file
	code := strings.Split(string(f), "\n")

	states := make(map[int]map[string]model.Code)

	// Insert turing code in memory (hashmap)
	for k, i := range code {
		line := strings.Split(i, " ")

		if len(line) == StandartLength {
			state, err := strconv.Atoi(line[0][1:])
			if err != nil {
				log.Fatalf("failed to read state in line: %d \t  %s", k+1, i)
			}

			nextState, err := strconv.Atoi(line[3][1:])
			if err != nil {
				log.Fatalf("failed to read next state in line: %d \t %s", k+1, i)
			}

			direction := string(line[5][0])
			lineCode := model.Code{
				CurrentState:  state,
				CurrentSymbol: line[1],
				NextState:     nextState,
				NextSymbol:    line[4],
				Direction:     direction,
			}

			// Validate code
			if e := lineCode.IsValid(); e != nil {
				log.Fatalf("there is error in code: %v", e)
			}

			if states[state] == nil {
				states[state] = make(map[string]model.Code)
			}

			states[state][line[1]] = lineCode
		}

	}

	t, err := os.ReadFile("tape.txt")
	if err != nil {
		log.Fatalf("failed to read tape.txt: %w", err)
	}

	// Read count of args
	tapeArgs := strings.Split(string(t), " ")

	countArgs, err := strconv.Atoi(tapeArgs[0])
	if err != nil {
		log.Fatalf("failed to read count of tape's arguments: %w", err)
	}

	// Create tape with zero value (ZeroCell)
	tape := model.NewTape(ZeroCell)

	start := tape

	for i := range countArgs {
		currentArg, err := strconv.Atoi(tapeArgs[i+1])
		if err != nil {
			log.Fatalf("failed to read tape's argument: %v", err)
		}

		for i := 0; i <= currentArg; i++ {
			tape = CreateCell(FullCell, tape)
		}

		tape = CreateCell(ZeroCell, tape)
	}
	fmt.Fprintf(writer, "Лента\n")
	TapeWrite(start, nil, writer)
	fmt.Fprintln(writer)

	leftCell := start.Next
	TapeWrite(start, leftCell, writer)
	currentState := states[1][leftCell.Symbol]
	for {
		currentState = states[currentState.CurrentState][leftCell.Symbol]
		if _, ok := states[currentState.NextState][leftCell.Symbol]; !ok && currentState.NextState != 0 {
			log.Fatalf("no code for this state: q%d %s", currentState.NextState, currentState.CurrentSymbol)
		}

		currentState = states[currentState.CurrentState][leftCell.Symbol]

		direction := currentState.Direction
		symbol := currentState.NextSymbol

		currentState = states[currentState.NextState][leftCell.Symbol]
		leftCell.Symbol = symbol

		switch direction {
		case "R":
			leftCell = leftCell.Next

		case "L":
			leftCell = leftCell.Last
		}

		TapeWrite(start, leftCell, writer)

		if currentState.CurrentState == 0 {
			break
		}

		GenerateEdge(leftCell)
	}

	newTape := SearchStart(start)

	fmt.Fprintf(writer, "\nNew tape\n")
	TapeWrite(newTape, nil, writer)

	writer.Flush()
}

func Contains(arr []string, element string) bool {
	for _, i := range arr {
		if element == i {
			return true
		}
	}

	return false
}

func SearchStart(tape *model.Tape) *model.Tape {
	if tape.Last != nil {
		tape = tape.Last
		SearchStart(tape)

	}
	return tape
}

func GenerateEdge(tape *model.Tape) *model.Tape {
	if tape.Last == nil {
		cell := model.Tape{
			Symbol: ZeroCell,
			Next:   tape,
		}

		tape.Last = &cell
	}
	if tape.Next == nil {
		cell := model.Tape{
			Symbol: ZeroCell,
			Last:   tape,
		}

		tape.Next = &cell
	}

	return tape
}

func ColorPaint(tape, current *model.Tape) bool {
	return tape == current
}

func TapeWrite(tape, current *model.Tape, writer *bufio.Writer) {
	for tape.Next != nil {
		if ColorPaint(tape, current) {
			fmt.Fprintf(writer, Red+"[%s]"+Reset, tape.Symbol)
		} else {
			fmt.Fprintf(writer, "[%s]", tape.Symbol)
		}

		tape = tape.Next
	}
	fmt.Fprintln(writer)
}

func CreateCell(symbol string, tape *model.Tape) *model.Tape {
	cell := model.Tape{
		Symbol: symbol,
		Last:   tape,
	}

	tape.Next = &cell
	tape = tape.Next

	return tape
}
