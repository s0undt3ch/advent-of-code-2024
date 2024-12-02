package matrix

import (
	"log"
	"strconv"
	"strings"
)

type Vector struct {
	Values []int
}

func ParseVector(line string) Vector {
	parts := strings.Split(line, " ")
	vector := make([]int, len(parts))
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			log.Fatalln(err)
		}
		vector[i] = n
	}
	return VectorFromSlice(vector)
}

func (v *Vector) All(f func(a int, b int) bool) bool {
	for i := 0; i < len(v.Values)-1; i++ {
		a := v.Values[i]
		b := v.Values[i+1]
		if !f(a, b) {
			return false
		}
	}
	return true
}

func (v *Vector) Any(f func(a int, b int) bool) bool {
	for i := 0; i < len(v.Values)-1; i++ {
		a := v.Values[i]
		b := v.Values[i+1]
		if f(a, b) {
			return true
		}
	}
	return false
}

func VectorFromSlice(s []int) Vector {
	return Vector{s}
}

func (v *Vector) Remove(s int) Vector {
	newValues := make([]int, 0, len(v.Values)-1)
	newValues = append(newValues, v.Values[:s]...)
	newValues = append(newValues, v.Values[s+1:]...)
    return Vector{newValues}
}
