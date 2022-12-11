package main

import "fmt"

const (
	CRT_WIDTH    = 40
	CRT_HEIGHT   = 6
	SPRITE_WIDTH = 3
)

type CrtImage [CRT_HEIGHT][CRT_WIDTH]byte

func parseCrtXs(input []string) []int {
	x := 1
	xs := []int{}
	var arg int
	for _, instr := range input {
		xs = append(xs, x)
		if _, err := fmt.Sscanf(instr, "addx %d", &arg); err == nil {
			xs = append(xs, x)
			x += arg
		} else if instr != "noop" {
			panic("Invalid input: " + instr)
		}
	}
	return xs
}

func signalStrengthSum(xs []int) int {
	sum := 0
	for _, i := range []int{20, 60, 100, 140, 180, 220} {
		sum += i * xs[i-1]
	}
	return sum
}

func renderCrtImage(xs []int) CrtImage {
	image := CrtImage{}
	xPos := 0
	yPos := 0
	for _, spriteXPos := range xs {
		if xPos == CRT_WIDTH {
			yPos++
			xPos = 0
		}
		if spriteXPos >= xPos-1 && spriteXPos < xPos+SPRITE_WIDTH-1 {
			image[yPos][xPos] = '#'
		} else {
			image[yPos][xPos] = '.'
		}
		xPos++
	}
	return image
}

func displayCrtImage(image CrtImage) {
	for _, line := range image {
		fmt.Println(string(line[:]))
	}
}

func day10(input []string) {
	crtXs := parseCrtXs(input)
	fmt.Println(signalStrengthSum(crtXs))
	displayCrtImage(renderCrtImage(crtXs))
}

func init() {
	Solutions[10] = day10
}

// 14600 is too high
