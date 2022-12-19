package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CrateStack []byte
type CrateStorage []CrateStack
type crateMoves struct {
	Cnt  int
	From int
	To   int
}

func (cStack *CrateStack) push(crate byte) {
	*cStack = append(*cStack, crate)
}

func (cStack *CrateStack) pop() byte {
	newLen := len(*cStack) - 1
	crate := (*cStack)[newLen]
	*cStack = (*cStack)[:newLen]
	return crate
}

func parseCrateStorage(input []string, stackDesc string) CrateStorage {
	stackDescVec := strings.Fields(stackDesc)
	numStacks, err := strconv.Atoi(stackDescVec[len(stackDescVec)-1])
	if err != nil {
		panic("Failed to figure out number of stacks")
	}
	cs := make(CrateStorage, numStacks)
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
		if _, err := fmt.Sscanf(row, "move %d from %d to %d", &moves.Cnt, &moves.From, &moves.To); err != nil {
			panic("Failed to parse move")
		}
		cm = append(cm, moves)
	}
	return cm
}

func parseCrates(input []string) (CrateStorage, []crateMoves) {
	for i, row := range input {
		if row == "" {
			return parseCrateStorage(input[:i-1], input[i-1]), parseCrateMoves(input[i+1:])
		}
	}
	panic("Invalid input")
}

func copyCrateStorage(cs CrateStorage) CrateStorage {
	copy := append(CrateStorage{}, cs...)
	for i := range copy {
		copy[i] = append(CrateStack{}, copy[i]...)
	}
	return copy
}

func createMsg(crates CrateStorage) string {
	msg := []byte{}
	for _, crate := range crates {
		msg = append(msg, crate.pop())
	}
	return string(msg)
}

func MoveCrates9000(crates CrateStorage, moveInstr []crateMoves) string {
	for _, moves := range moveInstr {
		for cnt := 0; cnt < moves.Cnt; cnt++ {
			crates[moves.To-1].push(crates[moves.From-1].pop())
		}
	}
	return createMsg(crates)
}

func MoveCrates9001(crates CrateStorage, moveInstr []crateMoves) string {
	stack := CrateStack{}
	for _, moves := range moveInstr {
		for cnt := 0; cnt < moves.Cnt; cnt++ {
			stack.push(crates[moves.From-1].pop())
		}
		for len(stack) > 0 {
			crates[moves.To-1].push(stack.pop())
		}
	}
	return createMsg(crates)
}

func day5(input []string) {
	storage, moves := parseCrates(input)
	fmt.Println(MoveCrates9000(copyCrateStorage(storage), moves))
	fmt.Println(MoveCrates9001(storage, moves))
}

func init() {
	Solutions[5] = day5
}
