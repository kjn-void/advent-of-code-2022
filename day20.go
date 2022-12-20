package main

import (
	"fmt"
	"strconv"
)

const DecryptionKey = 811589153

type Number struct {
	Value int64
	Prev  *Number
	Next  *Number
}

type EncryptedFile struct {
	Zero    *Number
	Numbers []*Number
}

func (number *Number) moveBackwards(period int) {
	steps := int(-number.Value % int64(period-1))
	for i := 0; i < steps; i++ {
		prev := number.Prev
		prev.Prev.Next, prev.Next, number.Next, prev.Prev, number.Prev, number.Next.Prev =
			number, number.Next, prev, number, prev.Prev, prev
	}
}

func (number *Number) moveForward(period int) {
	steps := int(number.Value % int64(period-1))
	for i := 0; i < steps; i++ {
		next := number.Next
		number.Prev.Next, number.Next, next.Next, number.Prev, next.Prev, next.Next.Prev =
			next, next.Next, number, next, number.Prev, number
	}
}

func (number *Number) move(period int) {
	if number.Value < 0 {
		number.moveBackwards(period)
	} else {
		number.moveForward(period)
	}
}

func mix(encryptedFile EncryptedFile) {
	for _, number := range encryptedFile.Numbers {
		number.move(len(encryptedFile.Numbers))
	}
}

func getNumbers(encryptedFile EncryptedFile) []int64 {
	number := encryptedFile.Zero
	originalFile := make([]int64, len(encryptedFile.Numbers))
	for i := 0; i < len(originalFile); i++ {
		originalFile[i] = number.Value
		number = number.Next
	}
	return originalFile
}

func calculateGroveCoordinates(numbers []int64) int64 {
	cnt := len(numbers)
	sum := numbers[1000%cnt]
	sum += numbers[2000%cnt]
	sum += numbers[3000%cnt]
	return sum
}

func GroveCoordinates(encryptedFile EncryptedFile) int64 {
	mix(encryptedFile)
	return calculateGroveCoordinates(getNumbers(encryptedFile))
}

func applyDecryptKey(encryptedFile EncryptedFile) {
	for _, number := range encryptedFile.Numbers {
		number.Value *= DecryptionKey
	}
}

func DecryptedGroveCoordinates(encryptedFile EncryptedFile) int64 {
	applyDecryptKey(encryptedFile)
	for i := 0; i < 10; i++ {
		mix(encryptedFile)
	}
	return calculateGroveCoordinates(getNumbers(encryptedFile))
}

func day20(input []string) {
	encryptedFile := parseEncryptedFile(input)
	fmt.Println(GroveCoordinates(encryptedFile))
	encryptedFile = parseEncryptedFile(input)
	fmt.Println(DecryptedGroveCoordinates(encryptedFile))
}

func init() {
	Solutions[20] = day20
}

func parseEncryptedFile(input []string) EncryptedFile {
	encryptedFile := EncryptedFile{Numbers: []*Number{}}
	var first, prev *Number
	for i, numStr := range input {
		number := Number{Prev: prev}
		if val, err := strconv.Atoi(numStr); err == nil {
			number.Value = int64(val)
		} else {
			panic("Failed to parse number: " + err.Error())
		}
		if number.Value == 0 {
			encryptedFile.Zero = &number
		}
		if i == 0 {
			first = &number
		} else {
			number.Prev = prev
			prev.Next = &number
		}
		prev = &number
		encryptedFile.Numbers = append(encryptedFile.Numbers, &number)
	}
	first.Prev = prev
	prev.Next = first
	return encryptedFile
}
