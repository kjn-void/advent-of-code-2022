package main

import (
	"fmt"
	"strconv"
	"strings"
)

type crateStack []byte
type crateStorage []crateStack
type crateMoves struct {
	cnt  int
	from int
	to   int
}

func (cStack *crateStack) push(crate byte) {
	*cStack = append(*cStack, crate)
}

func (cStack *crateStack) pop() byte {
	newLen := len(*cStack) - 1
	crate := (*cStack)[newLen]
	*cStack = (*cStack)[:newLen]
	return crate
}

func parseCrateStorage(input []string, stackDesc string) crateStorage {
	stackDescVec := strings.Fields(stackDesc)
	numStacks, err := strconv.Atoi(stackDescVec[len(stackDescVec)-1])
	if err != nil {
		panic("Failed to figure out number of stacks")
	}
	cs := make(crateStorage, numStacks)
	for row := len(input) - 1; row >= 0; row-- {
		line := input[row]
		for stack := 0; stack < numStacks; stack++ {
			crate := line[1+stack*4]
			if crate != ' ' {
				cs[stack].push(crate)
			}
		}
	}
	return cs
}

func parseCrateMoves(input []string) []crateMoves {
	cm := []crateMoves{}
	for _, row := range input {
		moves := crateMoves{}
		if _, err := fmt.Sscanf(row, "move %d from %d to %d", &moves.cnt, &moves.from, &moves.to); err != nil {
			panic("Failed to parse move")
		}
		cm = append(cm, moves)
	}
	return cm
}

func parseCrates(input []string) (crateStorage, []crateMoves) {
	for i, row := range input {
		if row == "" {
			return parseCrateStorage(input[:i-1], input[i-1]), parseCrateMoves(input[i+1:])
		}
	}
	panic("Invalid input")
}

func copyCrateStorage(cs crateStorage) crateStorage {
	copy := append(crateStorage{}, cs...)
	for i := range copy {
		copy[i] = append(crateStack{}, copy[i]...)
	}
	return copy
}

func createMsg(crates crateStorage) string {
	msg := []byte{}
	for _, crate := range crates {
		msg = append(msg, crate.pop())
	}
	return string(msg)
}

func moveCrates9000(crates crateStorage, moveInstr []crateMoves) string {
	for _, moves := range moveInstr {
		for cnt := 0; cnt < moves.cnt; cnt++ {
			crates[moves.to-1].push(crates[moves.from-1].pop())
		}
	}
	return createMsg(crates)
}

func moveCrates9001(crates crateStorage, moveInstr []crateMoves) string {
	stack := crateStack{}
	for _, moves := range moveInstr {
		for cnt := 0; cnt < moves.cnt; cnt++ {
			stack.push(crates[moves.from-1].pop())
		}
		for len(stack) > 0 {
			crates[moves.to-1].push(stack.pop())
		}
	}
	return createMsg(crates)
}

func day5(input []string) {
	storage, moves := parseCrates(input)
	fmt.Println(moveCrates9000(copyCrateStorage(storage), moves))
	fmt.Println(moveCrates9001(storage, moves))
}

func init() {
	Solutions[5] = day5
}
