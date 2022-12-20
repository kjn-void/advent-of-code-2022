package main

import (
	"testing"
)

var input20 = []string{"1", "2", "-3", "3", "-2", "0", "4"}

func TestDay20_1(t *testing.T) {
	encryptedFile := parseEncryptedFile(input20)
	if len(encryptedFile.Numbers) != len(input20) {
		t.Fatalf("Expected %d numbers, got %d", len(input20), len(encryptedFile.Numbers))
	}
	if encryptedFile.Zero.Prev.Value != -2 {
		t.Fatalf("Number before 0 should be -2, is %d", encryptedFile.Zero.Prev.Value)
	}
	if encryptedFile.Zero.Next.Value != 4 {
		t.Fatalf("Number after 0 should be 4, is %d", encryptedFile.Zero.Next.Value)
	}
}

func TestDay20_2(t *testing.T) {
	encryptedFile := parseEncryptedFile(input20)
	numbers := getNumbers(encryptedFile)
	expectedFile := []int64{0, 4, 1, 2, -3, 3, -2}
	if !cmpNumbers(expectedFile, numbers) {
		t.Fatalf("Number series differ, expected %v got %v", expectedFile, numbers)
	}
}

func TestDay20_3(t *testing.T) {
	encryptedFile := parseEncryptedFile(input20)
	encryptedFile.Numbers[0].move(len(encryptedFile.Numbers) - 1)
	numbers := getNumbers(encryptedFile)
	expectedFile := []int64{0, 4, 2, 1, -3, 3, -2}
	if !cmpNumbers(expectedFile, numbers) {
		t.Fatalf("Number series differ, expected %v got %v", expectedFile, numbers)
	}
}

func TestDay20_4(t *testing.T) {
	encryptedFile := parseEncryptedFile(input20)
	mix(encryptedFile)
	numbers := getNumbers(encryptedFile)
	expectedFile := []int64{0, 3, -2, 1, 2, -3, 4}
	if !cmpNumbers(expectedFile, numbers) {
		t.Fatalf("Number series differ, expected %v got %v", expectedFile, numbers)
	}
}

func TestDay20_5(t *testing.T) {
	encryptedFile := parseEncryptedFile(input20)
	grooveCoord := GroveCoordinates(encryptedFile)
	if grooveCoord != 3 {
		t.Fatalf("Expected groove coordinates to be 3, got %d", grooveCoord)
	}
}

func cmpNumbers(xs, ys []int64) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i := 0; i < len(xs); i++ {
		if xs[i] != ys[i] {
			return false
		}
	}
	return true
}

func BenchmarkDay20_parse(b *testing.B) {
	input := inputAsString(20)
	for n := 0; n < b.N; n++ {
		parseEncryptedFile(input)
	}
}

func BenchmarkDay20_part1(b *testing.B) {
	encryptedFile := parseEncryptedFile(inputAsString(20))
	for n := 0; n < b.N; n++ {
		GroveCoordinates(encryptedFile)
	}
}

func BenchmarkDay20_part2(b *testing.B) {
	input := inputAsString(20)
	for n := 0; n < b.N; n++ {
		DecryptedGroveCoordinates(parseEncryptedFile(input))
	}
}
