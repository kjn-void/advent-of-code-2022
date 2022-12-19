package main

import "fmt"

const (
	START_OF_PACKET_MARKER_LEN  = 4
	START_OF_MESSAGE_MARKER_LEN = 14
)

func isMarker(subSignal string) bool {
	seen := 0
	for _, ch := range subSignal {
		chBit := 1 << (ch - 'a')
		if seen&chBit != 0 {
			return false
		}
		seen |= chBit
	}
	return true
}

func markerOffset(signal string, markerLen int) int {
	for offset := 0; offset < len(signal)-markerLen-1; offset++ {
		maybeMarker := signal[offset : offset+markerLen]
		if isMarker(maybeMarker) {
			return offset + markerLen
		}
	}
	panic("No marker found")

}

func StartOfPacketMarker(signal string) int {
	return markerOffset(signal, START_OF_PACKET_MARKER_LEN)
}

func startOfMessageMarker(signal string) int {
	return markerOffset(signal, START_OF_MESSAGE_MARKER_LEN)
}

func day6(input []string) {
	signal := input[0]
	fmt.Println(StartOfPacketMarker(signal))
	fmt.Println(startOfMessageMarker(signal))
}

func init() {
	Solutions[6] = day6
}
