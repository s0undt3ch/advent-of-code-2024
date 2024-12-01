package day1

import (
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code-2024/internal/utils/file-reader"
)

func Main() {
	log.Println("DAY 1")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDay(1, 1)
	list1 := []int{}
	list2 := []int{}
	for f.HasMore() {
		line, _, err := f.Read()
		if err != nil {
			log.Fatalln(err)
		}
        log.Println(line)

		parts := strings.Split(line, "   ")
		number1, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		number2, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalln(err)
		}

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}
	sort.Sort(sort.IntSlice(list1))
	sort.Sort(sort.IntSlice(list2))

	sumDiffs := 0
	for i, l1 := range list1 {
		l2 := list2[i]

		diff := math.Abs(float64(l1 - l2))
		sumDiffs = sumDiffs + int(diff)
	}

	log.Println("The solution is:", sumDiffs)
}

func part2() {

}
