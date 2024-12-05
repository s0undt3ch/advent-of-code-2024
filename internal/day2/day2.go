package day2

import (
	"log"

	filereader "github.com/s0undt3ch/advent-of-code-2024/internal/utils/file-reader"
)

func Main() {
	log.Println("DAY 2")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayInput(2, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Lines:\n", lines)

}

func part2() {
	f := filereader.NewFromDayInput(2, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Lines:\n", lines)

}
