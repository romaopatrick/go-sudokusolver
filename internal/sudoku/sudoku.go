package sudoku

import "github.com/ledongthuc/goterators"

type Sudoku struct {
	Height              int
	Width               int
	BigSquareDefinition *BigSquareDefinition
	Board               [][]int
	hpn                 int
	possibleNums        []int
}

type BigSquareDefinition struct {
	Height int
	Width  int
}

func (b *BigSquareDefinition) GetHighestPossibleNumber() int {
	return b.Height * b.Width
}
func NewSudoku(
	bigSquareDefinition *BigSquareDefinition,
	initialBoard [][]int) *Sudoku {

	height := len(initialBoard)
	width := len(initialBoard[0])
	hpn := bigSquareDefinition.GetHighestPossibleNumber()
	possibleNums := make([]int, hpn-1)
	for i := 1; i <= hpn; i++ {
		possibleNums = append(possibleNums, i)
	}

	return &Sudoku{
		Height:              height,
		Width:               width,
		BigSquareDefinition: bigSquareDefinition,
		Board:               initialBoard,
		hpn:                 hpn,
		possibleNums:        possibleNums,
	}
}

func (s *Sudoku) Solve() [][]int {
	for x := range s.Board {
		for y := range s.Board[x] {
			elem := s.Board[x][y]
			if elem != 0 {
				continue
			}

			for i := 1; i <= s.hpn; i++ {
				colNums := s.columnPossibleNums(y)
				rowNums := s.rowPossibleNums(x)
				squareNums := s.squarePossibleNums()
			}
		}
	}

	return s.Board
}

func (s *Sudoku) columnPossibleNums(y int) []int {
	column := []int{}

	for _, row := range s.Board {
		elem := row[y]
		if elem > 0 {
			column = append(column, elem)
		}
	}

	result := goterators.Filter(s.possibleNums, func(i int) bool {
		return !goterators.Exist(column, i)
	})

	return result
}

func (s *Sudoku) rowPossibleNums(x int) []int {
	result := goterators.Filter(s.possibleNums, func(i int) bool {
		return !goterators.Exist(s.Board[x], i)
	})

	return result
}

func (s *Sudoku) squarePossibleNums(x, y int) []int {
	squareRows := []int{}
	squareCols := []int{}
	squareColsAmount := (s.Width * s.BigSquareDefinition.Width) / s.Width
	for idx, row := range s.Board {
		if x 
	}
}
