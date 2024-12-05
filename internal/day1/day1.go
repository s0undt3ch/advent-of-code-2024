package day1

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	filereader "github.com/s0undt3ch/advent-of-code-2024/internal/utils/file-reader"
)

func Main() {
	log.Println("DAY 1")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayInput(1, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Lines:\n", lines)

	var left []int
	var right []int
	var diff int

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) <= 0 {
			continue
		}
		result := strings.Fields(lines[i])
		log.Println("Result:", result)
		il, err := strconv.Atoi(result[0])
		if err != nil {
			panic(err)
		}
		left = append(left, il)

		ir, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}
		right = append(right, ir)
	}

	log.Println("Lines split, slices fed")

	// Sort left
	sort.Ints(left)
	sort.Ints(right)

	log.Println("Left:", left)
	log.Println("Right:", right)

	for i := 0; i < len(left); i++ {
		sum := math.Abs(float64(left[i] - right[i]))
		diff = diff + int(sum)
	}

	log.Println("Diff Slice:", diff)
}

func part2() {
	f := filereader.NewFromDayInput(1, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Lines:\n", lines)

	var left []int

	counts := make(map[int]int)
	similarity := 0

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) <= 0 {
			continue
		}
		result := strings.Fields(lines[i])
		log.Println("Result:", result)
		il, err := strconv.Atoi(result[0])
		if err != nil {
			panic(err)
		}
		left = append(left, il)

		ir, err := strconv.Atoi(result[1])
		if err != nil {
			panic(err)
		}

		current_count := counts[ir]
		counts[ir] = current_count + 1
	}

	log.Println("Lines split, counts done")
	log.Println("Counts:")
	for key, value := range counts {
		log.Println("Key:", key, "Value:", value)
	}

	for idx := range left {
		number := left[idx]
		count := counts[number]
		log.Println("Number:", number, "Count:", count)
		similarity = similarity + (number * count)
	}

	log.Println("Similarity:", similarity)
}
