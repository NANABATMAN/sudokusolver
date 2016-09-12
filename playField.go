package main

import (
	"fmt"
)

const (
	SIZE = 9
)

type PlayField struct {
	Matrix [SIZE][SIZE]int
}

func (playField *PlayField) SetPlayField(matrix *[SIZE][SIZE]int) {
	playField.Matrix = (*matrix)
}

func (playField *PlayField) GetPlayField() *[SIZE][SIZE]int {
	return &playField.Matrix
}

func (playField *PlayField) Solve() bool {
	nextRowIdx, nextColIdx, hasEmptyCell := playField.findEmptyCell()

	if !hasEmptyCell {
		return true
	}

	for number := 1; number <= SIZE; number++ {
		if playField.isAllowedNumber(nextRowIdx, nextColIdx, number) {
			playField.Matrix[nextRowIdx][nextColIdx] = number

			if playField.Solve() {
				return true
			}

			playField.Matrix[nextRowIdx][nextColIdx] = 0
		}
	}

	return false
}

func (playField *PlayField) findEmptyCell() (int, int, bool) {
	for rowIdx := 0; rowIdx < SIZE; rowIdx++ {
		for colIdx := 0; colIdx < SIZE; colIdx++ {
			if playField.Matrix[rowIdx][colIdx] == 0 {
				return rowIdx, colIdx, true
			}
		}
	}

	return 0, 0, false
}

func (playField *PlayField) isAllowedNumber(rowIdx, colIdx, number int) bool {
	subRowIdxStart := getIndexOfSubField(rowIdx)
	subColIdxStart := getIndexOfSubField(colIdx)

	for i := 0; i < SIZE; i++ {
		if playField.Matrix[rowIdx][i] == number || playField.Matrix[i][colIdx] == number ||
			playField.Matrix[subRowIdxStart+i/3][subColIdxStart+i%3] == number {
			return false
		}
	}

	return true
}

func (playField *PlayField) PrintPlayField() {
	for _, row := range playField.Matrix {
		for _, el := range row {
			fmt.Print(el, " ")
		}
		fmt.Println()
	}
}

func getIndexOfSubField(index int) int {
	if index < 3 {
		return 0
	} else if index < 6 {
		return 3
	} else {
		return 6
	}
}
