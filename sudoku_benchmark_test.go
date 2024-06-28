package gosudokusolver_test

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"

	sudoku "github.com/romaopatrick/go-sudokusolver"
)

type testData struct {
	input,
	output []int
}

var MAXLINES = 50020

func BenchmarkSolve(b *testing.B) {
	f, _ := os.Open("./tests/dataset.csv")
	r := bufio.NewReader(f)
	currentLine := 0
	data := []testData{}

	for currentLine < MAXLINES {
		l, _, errS := r.ReadLine()
		if errS != nil {
			break
		}

		lStr := string(l)
		inOutStr := strings.Split(lStr, ",")
		inStr := inOutStr[0]

		inSplit := strings.Split(inStr, "")
		input := make([]int, len(inSplit))
		for i, c := range inSplit {
			input[i], _ = strconv.Atoi(c)
		}

		outStr := inOutStr[1]
		outSplit := strings.Split(outStr, "")
		output := make([]int, len(outSplit))
		for i, c := range outSplit {
			output[i], _ = strconv.Atoi(c)
		}

		dataSet := testData{
			input,
			output,
		}

		data = append(data, dataSet)

		currentLine++
	}

	b.ReportAllocs()
	b.ResetTimer()
	for _, d := range data {
		s := sudoku.NewSudoku(9, 9, 3, 3, d.input, false)
		s.Solve()
	}
}
