package main

import (
	"fmt"
	"sort"
	"strconv"
	"unicode"
)

const (
	RIGHT_ORDER PacketOrder = iota
	WRONG_ORDER
	EQUAL_ORDER
)

type PacketOrder uint8

type PacketValue struct {
	list   []PacketValue
	number int
}

type Packet struct {
	left  PacketValue
	right PacketValue
}

func numberToList(number int) []PacketValue {
	value := PacketValue{}
	value.number = number
	return []PacketValue{value}
}

func areListsInRightOrder(left, right []PacketValue) PacketOrder {
	for i := 0; ; i++ {
		if len(left) == i {
			if len(right) == i {
				return EQUAL_ORDER
			}
			return RIGHT_ORDER
		}
		if len(right) == i {
			return WRONG_ORDER
		}
		order := areValuesInRightOrder(left[i], right[i])
		if order != EQUAL_ORDER {
			return order
		}
	}
}

func areValuesInRightOrder(left, right PacketValue) PacketOrder {
	if left.list == nil {
		if right.list == nil {
			switch {
			case left.number < right.number:
				return RIGHT_ORDER
			case left.number > right.number:
				return WRONG_ORDER
			default:
				return EQUAL_ORDER
			}
		}
		return areListsInRightOrder(numberToList(left.number), right.list)
	} else {
		if right.list == nil {
			return areListsInRightOrder(left.list, numberToList(right.number))
		}
		return areListsInRightOrder(left.list, right.list)
	}
}

func (packet Packet) isInRightOrder() bool {
	return areValuesInRightOrder(packet.left, packet.right) == RIGHT_ORDER
}

func SumPacketIndicesInRightOrder(packets []Packet) int {
	sum := 0
	for i, packet := range packets {
		if packet.isInRightOrder() {
			sum += i + 1
		}
	}
	return sum
}

func packetList(packets []Packet) []PacketValue {
	list := make([]PacketValue, len(packets)*2)
	for i, p := range packets {
		list[i*2] = p.left
		list[i*2+1] = p.right
	}
	return list
}

func DividerPackets() [2]PacketValue {
	divider := [2]PacketValue{}
	col := 0
	divider[0] = parsePacketValue("[[2]]", &col)
	col = 0
	divider[1] = parsePacketValue("[[6]]", &col)
	return divider
}

func DecoderKey(packets []PacketValue) int {
	divider := DividerPackets()
	packets = append(packets, divider[:]...)

	sort.Slice(packets, func(i, j int) bool { return areValuesInRightOrder(packets[i], packets[j]) == RIGHT_ORDER })

	key := 1
	d := 0
	for i, packet := range packets {
		if areValuesInRightOrder(packet, divider[d]) == EQUAL_ORDER {
			d++
			key *= (i + 1)
		}
		if d == len(divider) {
			break
		}
	}
	return key
}

func day13(input []string) {
	packets := parsePackets(input)
	fmt.Println(SumPacketIndicesInRightOrder(packets))
	fmt.Println(DecoderKey(packetList(packets)))
}

func init() {
	Solutions[13] = day13
}

func parsePacketNumber(value string, col *int) int {
	c := *col
	for unicode.IsNumber(rune(value[*col])) {
		*col += 1
	}
	if number, err := strconv.Atoi(value[c:*col]); err == nil {
		return number
	}
	return -1
}

func parsePacketList(value string, col *int) []PacketValue {
	list := []PacketValue{}
	for value[*col] != ']' {
		*col += 1
		packetValue := parsePacketValue(value, col)
		if packetValue.list != nil || packetValue.number >= 0 {
			list = append(list, packetValue)
		}
	}
	*col += 1
	return list
}

func parsePacketValue(value string, col *int) PacketValue {
	packetValue := PacketValue{}
	if value[*col] == '[' {
		packetValue.number = -1
		packetValue.list = parsePacketList(value, col)
	} else {
		packetValue.number = parsePacketNumber(value, col)
	}
	return packetValue
}

func parsePacket(input []string, row *int) Packet {
	packet := Packet{}
	col := 0
	packet.left = parsePacketValue(input[*row], &col)
	col = 0
	packet.right = parsePacketValue(input[*row+1], &col)
	*row += 3
	return packet
}

func parsePackets(input []string) []Packet {
	packets := []Packet{}
	row := 0
	for row < len(input) {
		packets = append(packets, parsePacket(input, &row))
	}
	return packets
}
