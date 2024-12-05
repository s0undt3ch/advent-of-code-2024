package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/s0undt3ch/advent-of-code-2024/internal/day1"
	"github.com/s0undt3ch/advent-of-code-2024/internal/day2"
	newday "github.com/s0undt3ch/advent-of-code-2024/internal/utils/new-day"
)

var days = map[int]func(){
	1: day1.Main,
	2: day2.Main,
}

func main() {
	log.SetFlags(0)

	fmt.Println("Advent of Code 2024")
	fmt.Println("Initial repo skeleton - All credits go to Afonso Raposo")
	fmt.Println("Also, this is my first stab at Go, as you might notice as you read through")

	day := flag.Int("day", 0, "Advent Day")
	newDay := flag.Bool("new-day", false, "Create a new day instead of running the passed day")

	flag.Parse()

	if *day <= 0 {
		log.Fatalln("The day must be > 1")
	}

	if dayFunc, ok := days[*day]; ok {
		dayFunc()
	} else if *newDay {
		text, render_err := newday.Render(*day)
		if render_err != nil {
			log.Fatal(render_err)
		}
		log.Println(text)
		new_day_path := filepath.Join(".", "internal", fmt.Sprintf("day%d", *day))
		mkdir_err := os.MkdirAll(new_day_path, os.ModePerm)
		if mkdir_err != nil {
			log.Fatal(mkdir_err)
		}
		new_day_file_path := filepath.Join(new_day_path, fmt.Sprintf("day%d.go", *day))
		f, open_file_err := os.OpenFile(new_day_file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if open_file_err != nil {
			log.Fatal(open_file_err)
		}
		_, write_err := f.WriteString(text)
		if write_err != nil {
			log.Fatal(write_err)
		}
	} else {
		log.Fatalln("Day not found")
	}

}
