package main

import "testing"

var input13 = []string{
	"[1,1,3,1,1]",
	"[1,1,5,1,1]",
	"",
	"[[1],[2,3,4]]",
	"[[1],4]",
	"",
	"[9]",
	"[[8,7,6]]",
	"",
	"[[4,4],4,4]",
	"[[4,4],4,4,4]",
	"",
	"[7,7,7,7]",
	"[7,7,7]",
	"",
	"[]",
	"[3]",
	"",
	"[[[]]]",
	"[[]]",
	"",
	"[1,[2,[3,[4,[5,6,7]]]],8,9]",
	"[1,[2,[3,[4,[5,6,0]]]],8,9]",
}

func TestDay13_1(t *testing.T) {
	packets := parsePackets(input13)
	if len(packets) != 8 {
		t.Fatalf("Expected 8 packets, got %d", len(packets))
	}
	if len(packets[0].left.list) != 5 {
		t.Fatalf("Expected left list to be of length 5, is %d", len(packets[0].left.list))
	}
	if packets[0].left.list[0].number != 1 || packets[0].left.list[2].number != 3 {
		t.Fatal("Incorrect content of packet")
	}
}

func TestDay13_2(t *testing.T) {
	packets := parsePackets(input13)
	sum := SumPacketIndicesInRightOrder(packets)
	if sum != 13 {
		t.Fatalf("Expected a sum of 13, got %d", sum)
	}
}

func TestDay13_3(t *testing.T) {
	packets := parsePackets(input13)
	decoderKey := DecoderKey(packetList(packets))
	if decoderKey != 140 {
		t.Fatalf("Decoder key should be 140, is %d", decoderKey)
	}
}

func BenchmarkDay13_parsing(b *testing.B) {
	input := inputAsString(13)
	for n := 0; n < b.N; n++ {
		parsePackets(input)
	}
}

func BenchmarkDay13_part1(b *testing.B) {
	packets := parsePackets(inputAsString(13))
	for n := 0; n < b.N; n++ {
		SumPacketIndicesInRightOrder(packets)
	}
}

func BenchmarkDay13_part2(b *testing.B) {
	packets := parsePackets(inputAsString(13))
	for n := 0; n < b.N; n++ {
		DecoderKey(packetList(packets))
	}
}
