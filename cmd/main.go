package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/s0undt3ch/advent-of-code-2024/internal/day1"
)

var days = map[int]func(){
	1: day1.Main,
}

func main() {
	log.SetFlags(0)

	fmt.Println("Advent of Code 2024")
	fmt.Println("Initial repo skeleton - All credits go to Afonso Raposo")
	fmt.Println("Also, this is my first stab at Go, as you might notice as you read through")

	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatalln("You must specify which day you want to run, e.g.: 1")
	}

	dayStr := args[0]

	day, err := strconv.Atoi(dayStr)
	if err != nil {
		log.Fatalln(err)
	}

	days[day]()
}
