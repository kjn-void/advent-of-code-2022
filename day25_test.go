package main

import "testing"

var input25 = []string{
	"1=-0-2",
	"12111",
	"2=0=",
	"21",
	"2=01",
	"111",
	"20012",
	"112",
	"1=-1=",
	"1-12",
	"12",
	"1=",
	"122",
}

func TestDay25_1(t *testing.T) {
	decimals := []int{
		1747,
		906,
		198,
		11,
		201,
		31,
		1257,
		32,
		353,
		107,
		7,
		3,
		37,
	}
	for i, snafu := range input25 {
		decimal := decimals[i]
		if snafuToDec(snafu) != decimal {
			t.Fatalf("Expected %s to translate to %d, got %d", snafu, decimal, snafuToDec(snafu))
		}
	}
}

func TestDay25_2(t *testing.T) {
	sum := 0
	for _, snafu := range input25 {
		sum += snafuToDec(snafu)
	}
	if sum != 4890 {
		t.Fatalf("Sum of SNAFU numbers should be 4890, got %d", sum)
	}
}

func TestDay25_3(t *testing.T) {
	if decToSnafu(4890) != "2=-1=0" {
		t.Fatalf("Converting 4890 to SNAFU should result in 2=-1=0, got %s", decToSnafu(4890))
	}
}

func BenchmarkDay25(b *testing.B) {
	input := inputAsString(25)
	for n := 0; n < b.N; n++ {
		ConsoleSnafu(input)
	}
}
