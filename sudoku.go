package gosudokusolver

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/ledongthuc/goterators"
)

type Sudoku struct {
	Height       int
	Width        int
	SquareHeight int
	SquareWidth  int

	IterationsToSolve int
	Input             []int
	Current           []int
	DebugMode         bool
}

func NewSudoku(
	h, w, sh, sw int,
	in []int,
	debugMode bool,
) *Sudoku {
	return &Sudoku{
		Height:       9,
		Width:        9,
		SquareHeight: 3,
		SquareWidth:  3,
		Input:        in,
		Current:      in,
		DebugMode:    debugMode,
	}
}
func (s *Sudoku) String() string {
	if !s.DebugMode {
		return ""
	}

	b := strings.Builder{}
	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("CURRENT ITERATION: %v \n", s.IterationsToSolve))
	countTwentySeven := 1
	countNine := 1
	countTree := 1
	for _, v := range s.Current {
		b.WriteString(fmt.Sprint(v))

		if countTree == 3 {
			countTree = 0
			b.WriteString("\t")
		}

		if countNine == 9 {
			countNine = 0
			b.WriteString("\n")
		}

		if countTwentySeven == 27 {
			countTwentySeven = 0
			b.WriteString("\n")
		}

		countTree++
		countNine++
		countTwentySeven++
	}

	return b.String()
}
func (s *Sudoku) Iterate() bool {
	for i := 0; i < len(s.Current); i++ {
		if s.Current[i] == 0 {
			for num := 1; num <= s.Height; num++ {
				if s.IsSafe(i, num) {
					s.Current[i] = num
					if s.Iterate() {
						return true
					}
					s.Current[i] = 0
				}

			}
			return false
		}
		if s.DebugMode {
			log.Println(s)
			log.Printf("%v to go! :)", goterators.Count(s.Current, 0))
			s.IterationsToSolve++
		}
	}
	return true
}

func (s *Sudoku) Solve() error {
	if !s.Iterate() {
		if s.DebugMode {
			log.Println(s)
			log.Printf("%v to go! :)", goterators.Count(s.Current, 0))
		}
		return errors.New("SOLVER BROKEN")
	}

	if s.DebugMode {
		log.Println(s)
		log.Printf("%v to go! :)", goterators.Count(s.Current, 0))
		log.Printf("TOTAL ITERATIONS TO SOLVE: %v", s.IterationsToSolve)
	}

	return nil
}

func (s *Sudoku) IsSafe(numIdx, num int) bool {
	return s.IsSafeCol(numIdx, num) &&
		s.IsSafeRow(numIdx, num) &&
		s.IsSafeSquare(numIdx, num)
}

func (s *Sudoku) IsSafeSquare(numIdx, num int) bool {
	idxCol := numIdx % s.Width
	idxRow := numIdx / s.Height

	sqIdx := (idxCol / s.SquareWidth) +
		(idxRow/s.SquareHeight)*
			(s.Width/s.SquareWidth)*
			s.SquareHeight

	sIdx := sqIdx * s.SquareWidth

	for line := 0; line < s.SquareHeight; line++ {
		for col := 0; col < s.SquareWidth; col++ {
			elem := s.Current[sIdx+col+(line*s.Width)]
			if elem == num {
				return false
			}
		}
	}

	return true
}

func (s *Sudoku) IsSafeCol(numIdx, num int) bool {

	var (
		sIdx int
	)
	if numIdx < s.Width {
		sIdx = numIdx
	} else {
		sIdx = (numIdx % s.Width)
	}

	for i := 0; i < s.Height; i++ {
		elem := s.Current[sIdx+(i*s.Width)]
		if elem == num {
			return false
		}
	}

	return true
}

func (s *Sudoku) IsSafeRow(numIdx, num int) bool {
	diff := float64(numIdx) / float64(s.Width)
	sIdx := int(math.Floor(diff)) * s.Width

	for i := 0; i < s.Width; i++ {
		elem := s.Current[sIdx+i]
		if elem == num {
			return false
		}
	}

	return true
}
