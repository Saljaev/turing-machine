package app

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"turing-machine/internal/config"
	"turing-machine/internal/logic"
	"turing-machine/internal/model"
)

func Run() {
	cfg := config.ConfigLoad()

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

		if len(line) == cfg.DefaultCodeLength {
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
		} else {
			log.Fatalf("there is no default code length in line: %d", k+1)
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
	tape := model.NewTape(cfg.ZeroCell)

	start := tape

	for i := range countArgs {
		currentArg, err := strconv.Atoi(tapeArgs[i+1])
		if err != nil {
			log.Fatalf("failed to read tape's argument: %v", err)
		}

		// Fill tape with args value
		for i := 0; i <= currentArg; i++ {
			tape = logic.CreateCell(cfg.FullCell, tape)
		}

		// Add separatin symbol between args
		tape = logic.CreateCell(cfg.ZeroCell, tape)
	}

	fmt.Fprintf(writer, "Лента\n")
	logic.TapeWrite(start, nil, writer)
	fmt.Fprintln(writer)

	leftCell := start.Next
	logic.TapeWrite(start, leftCell, writer)
	currentState := states[1][leftCell.Symbol].CurrentState

	// Execution all turing code
	for {
		leftCell, currentState, err = logic.Execute(states, currentState, leftCell)
		if err != nil {
			log.Fatalf("%v", err)
		}

		logic.TapeWrite(start, leftCell, writer)

		if currentState == 0 {
			break
		}

		logic.GenerateEdge(leftCell, cfg.ZeroCell)
	}

	newTape := logic.SearchStart(start)

	fmt.Fprintf(writer, "\nNew tape\n")
	logic.TapeWrite(newTape, nil, writer)

	writer.Flush()
}
