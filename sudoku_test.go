package sudoku_test

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"testing"

	sudoku "github.com/romaopatrick/go-sudokusolver"

	"github.com/ledongthuc/goterators"
	"github.com/stretchr/testify/assert"
)

type testInput struct {
	Input []int
}

const testJsonPath = "./tests/test.json"
const testInputPath = "./tests/input.txt"

func TestGetColumn(t *testing.T) {
	f, _ := os.ReadFile(testJsonPath)
	tIn := &testInput{}
	json.Unmarshal(f, tIn)

	inputs := []struct {
		in int
		e  []int
	}{
		{0, []int{0, 0, 8, 0, 3, 0, 0, 2, 7}},
		{1, []int{7, 4, 0, 9, 5, 0, 0, 1, 0}},
		{17, []int{3, 0, 0, 0, 0, 8, 0, 0, 0}},
		{40, []int{0, 0, 3, 5, 6, 0, 0, 0, 8}},
		{50, []int{0, 9, 4, 2, 0, 5, 9, 0, 0}},
	}

	s := sudoku.NewSudoku(9, 9, 3, 3, tIn.Input)
	a := assert.New(t)

	for _, input := range inputs {
		g := s.GetColumn(input.in)

		a.EqualValues(input.e, g)
	}
}

func TestGetRow(t *testing.T) {
	f, _ := os.ReadFile(testJsonPath)
	tIn := &testInput{}
	json.Unmarshal(f, tIn)

	inputs := []struct {
		in int
		e  []int
	}{
		{0, []int{0, 7, 0, 0, 0, 0, 0, 4, 3}},
		{2, []int{0, 7, 0, 0, 0, 0, 0, 4, 3}},
		{17, []int{0, 4, 0, 0, 0, 9, 6, 1, 0}},
		{80, []int{7, 0, 4, 0, 8, 0, 2, 0, 0}},
	}

	s := sudoku.NewSudoku(9, 9, 3, 3, tIn.Input)

	a := assert.New(t)
	for _, input := range inputs {
		g := s.GetRow(input.in)

		a.EqualValues(input.e, g)
	}
}

func TestCountSquares(t *testing.T) {
	f, _ := os.ReadFile(testJsonPath)
	tIn := &testInput{}
	json.Unmarshal(f, tIn)

	s := sudoku.NewSudoku(9, 9, 3, 3, tIn.Input)
	a := assert.New(t)

	g := s.CountSquares()
	e := 9

	a.EqualValues(e, g)
}

func TestGetSquare(t *testing.T) {
	f, _ := os.ReadFile(testJsonPath)
	tIn := &testInput{}
	json.Unmarshal(f, tIn)

	inputs := []struct {
		in int
		e  []int
	}{
		{0, []int{0, 7, 0, 0, 4, 0, 8, 0, 0}},
		{1, []int{0, 7, 0, 0, 4, 0, 8, 0, 0}},
		{2, []int{0, 7, 0, 0, 4, 0, 8, 0, 0}},
		{9, []int{0, 7, 0, 0, 4, 0, 8, 0, 0}},
		{10, []int{0, 7, 0, 0, 4, 0, 8, 0, 0}},
		{11, []int{0, 7, 0, 0, 4, 0, 8, 0, 0}},
		{18, []int{0, 7, 0, 0, 4, 0, 8, 0, 0}},
		{19, []int{0, 7, 0, 0, 4, 0, 8, 0, 0}},
		{20, []int{0, 7, 0, 0, 4, 0, 8, 0, 0}},

		{4, []int{0, 0, 0, 0, 0, 9, 6, 3, 4}},
		{13, []int{0, 0, 0, 0, 0, 9, 6, 3, 4}},

		{37, []int{0, 9, 4, 3, 5, 8, 0, 0, 8}},
	}

	s := sudoku.NewSudoku(9, 9, 3, 3, tIn.Input)

	a := assert.New(t)

	for _, input := range inputs {
		g := s.GetSquare(input.in)

		a.EqualValues(input.e, g)
	}
}

func TestSolve(t *testing.T) {
	f, _ := os.ReadFile(testInputPath)

	str := string(f)
	strInOut := strings.Split(str, "\n")

	inStr := strInOut[0]
	inSplit := strings.Split(inStr, "")
	in := make([]int, len(inSplit))
	for i, c := range inSplit {
		in[i], _ = strconv.Atoi(c)
	}

	outStr := strInOut[1]
	outSplit := strings.Split(outStr, "")
	out := make([]int, len(outSplit))
	for i, c := range outSplit {
		out[i], _ = strconv.Atoi(c)
	}

	s := sudoku.NewSudoku(9, 9, 3, 3, in)
	a := assert.New(t)

	err := s.Solve()
	a.Nil(err)
	re := goterators.Count(s.Current, 0)
	a.EqualValues(0, re)
	a.EqualValues(out, s.Current)
}
