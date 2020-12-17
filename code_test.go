package advent_test

import (
	advent "advent20201204"
	"fmt"
	"testing"
)

func TestRecordFromString(t *testing.T) {
	result := advent.RecordFromString("hgt:171cm hcl:#cfa07d pid:674448249")
	fmt.Println(result)
}
func TestRecordIterator(t *testing.T) {
	result := advent.RecordIterator("input.txt")
	fmt.Println(result)
}

func TestPart1(t *testing.T) {
	// fmt.Println(advent.Part1("input.txt"))
}
func TestPart2(t *testing.T) {
	fmt.Println(advent.Part2("input.txt"))
}
