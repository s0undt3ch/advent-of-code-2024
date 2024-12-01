package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
    "github.com/afonsocraposo/advent-of-code-2024/internal/day1"
    "github.com/afonsocraposo/advent-of-code-2024/internal/day2"
)


var days = map[int]func(){
    1:  day1.Main,
    2:  day2.Main,
}

func main() {
    fmt.Println("Advent of Code 2024")

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
