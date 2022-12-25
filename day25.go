package main

import "fmt"

func pow(n, e int) int {
	p := 1
	for i := 0; i < e; i++ {
		p *= n
	}
	return p
}

func snafuToDec(snafu string) int {
	values := map[rune]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}
	decimal := 0
	for i, snafuDigit := range snafu {
		decimal += values[snafuDigit] * pow(5, len(snafu)-i-1)
	}
	return decimal
}

func decToSnafu(n int) string {
	snafus := []string{"0", "1", "2", "=", "-"}
	snafu := ""
	for n > 0 {
		rem := n % 5
		snafu += snafus[rem]
		n /= 5
		if rem > 2 {
			n++
		}
	}
	return reverse(snafu)
}

func reverse(str string) string {
	revString := ""
	for _, v := range str {
		revString = string(v) + revString
	}
	return revString
}

func ConsoleSnafu(input []string) string {
	sum := 0
	for _, snafu := range input {
		sum += snafuToDec(snafu)
	}
	return decToSnafu(sum)
}

func day25(input []string) {
	fmt.Println(ConsoleSnafu(input))
}

func init() {
	Solutions[25] = day25
}
