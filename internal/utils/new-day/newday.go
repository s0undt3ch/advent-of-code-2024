package newday

import (
	"bytes"
	"text/template"
)

var newDayTemplate = `
package day{{ .DAY }}

import (
	"log"

	filereader "github.com/s0undt3ch/advent-of-code-2024/internal/utils/file-reader"
)

func Main() {
	log.Println("DAY {{ .DAY }}")

	log.Println("Part 1:")
	part1()

	log.Println("Part 2:")
	part2()
}

func part1() {
	f := filereader.NewFromDayInput({{ .DAY }}, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Lines:\n", lines)

}

func part2() {
	f := filereader.NewFromDayInput({{ .DAY }}, 1)

	lines, err := f.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Lines:\n", lines)
	
}
`

func Render(day int) (string, error) {
	ret, err := render_new_day_module(
		newDayTemplate,
		template.FuncMap{
			"DAY": day,
		},
	)
	return ret, err
}

func render_new_day_module(templateText string, funcMap template.FuncMap) (string, error) {
	var template = template.New("genericTemplate")
	t, err := template.Parse(templateText)
	if err != nil {
		return "", err
	}
	var buff bytes.Buffer
	execute_err := t.Execute(&buff, funcMap)
	if execute_err != nil {
		return "", execute_err
	}
	return buff.String(), nil
}
