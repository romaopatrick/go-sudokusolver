package gosudokusolver_test

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
const testInputPath = "./tests/input_3.txt"

func TestIsSafeCol(t *testing.T) {
	f, _ := os.ReadFile(testJsonPath)
	tIn := &testInput{}
	json.Unmarshal(f, tIn)

	inputs := []struct {
		inIdx int
		inNum int
		e     bool
	}{
		{0, 8, false},
		{1, 2, true},
		{17, 3, false},
		{40, 6, false},
		{50, 9, false},
	}

	s := sudoku.NewSudoku(9, 9, 3, 3, tIn.Input, true)
	a := assert.New(t)

	for idx, input := range inputs {
		if !a.EqualValues(input.e, s.IsSafeCol(input.inIdx, input.inNum)) {
			a.FailNowf("failed on: ", "%v idx", idx)
		}
	}
}

func TestIsSafeRow(t *testing.T) {
	f, _ := os.ReadFile(testJsonPath)
	tIn := &testInput{}
	json.Unmarshal(f, tIn)

	inputs := []struct {
		inIdx int
		inNum int
		e     bool
	}{
		{0, 7, false},
		{0, 1, true},
		{2, 4, false},
		{2, 8, true},
		{17, 9, false},
		{80, 1, true},
	}

	s := sudoku.NewSudoku(9, 9, 3, 3, tIn.Input, true)

	a := assert.New(t)
	for idx, input := range inputs {
		if !a.EqualValues(input.e, s.IsSafeRow(input.inIdx, input.inNum)) {
			a.FailNowf("failed on: ", "%v idx", idx)
		}
	}
}

func TestIsSafeSquare(t *testing.T) {
	f, _ := os.ReadFile(testJsonPath)
	tIn := &testInput{}
	json.Unmarshal(f, tIn)

	inputs := []struct {
		inIdx int
		inNum int
		e     bool
	}{
		{0, 7, false},
		{1, 1, true},
		{2, 8, false},
		{9, 9, true},
		{10, 6, true},
		{11, 6, true},
		{18, 8, false},
		{19, 3, true},
		{20, 4, false},
		{4, 6, false},
		{13, 1, true},
		{37, 8, false},
	}

	s := sudoku.NewSudoku(9, 9, 3, 3, tIn.Input, true)

	a := assert.New(t)

	for idx, input := range inputs {
		if !a.EqualValues(input.e, s.IsSafeSquare(input.inIdx, input.inNum)) {
			a.FailNowf("failed on: ", "%v idx", idx)
		}
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

	s := sudoku.NewSudoku(9, 9, 3, 3, in, true)
	a := assert.New(t)

	err := s.Solve()
	a.Nil(err)
	re := goterators.Count(s.Current, 0)
	a.EqualValues(0, re)
	a.EqualValues(out, s.Current)
}
