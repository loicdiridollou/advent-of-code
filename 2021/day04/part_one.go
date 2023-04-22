package main

import (
	"regexp"
	"strconv"
	"strings"
)

func StrToInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

type Board struct {
	board [][]string
}

func (board *Board) Play(num string) {
	for i := 0; i < len(board.board); i++ {
		for j := 0; j < len(board.board[0]); j++ {
			if board.board[i][j] == num {
				board.board[i][j] = "P"
			}
		}
	}
}

func (board *Board) Sum() int {
	val := 0
	for i := 0; i < len(board.board); i++ {
		for j := 0; j < len(board.board[0]); j++ {
			if board.board[i][j] != "P" {
				val += StrToInt(board.board[i][j])
			}
		}
	}
	return val
}

func (board *Board) HasWon() (bool, int) {
	winning := false
	for i := 0; i < len(board.board); i++ {
		numP := 0
		for j := 0; j < len(board.board[0]); j++ {
			if board.board[j][i] == "P" {
				numP++
			}
		}
		if numP == 5 {
			winning = true
			return true, board.Sum()
		}
	}
	for i := 0; i < len(board.board); i++ {
		numP := 0
		for j := 0; j < len(board.board[0]); j++ {
			if board.board[i][j] == "P" {
				numP++
			}
		}
		if numP == 5 {
			winning = true
			return true, board.Sum()
		}
	}
	return winning, 0
}

func prepareData(input []string) []Board {
	r, _ := regexp.Compile(`(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)`)
	lst_board := []Board{}
	for _, grp := range input {
		board := Board{make([][]string, 0)}
		for _, line := range strings.Split(grp, "\n") {
			if len(line) == 0 {
				continue
			}
			matches := r.FindStringSubmatch(line)
			board.board = append(board.board, matches[1:])
		}
		lst_board = append(lst_board, board)
	}
	return lst_board
}

func part1(input string) int {
	s := strings.Split(input, "\n\n")

	numbers := s[0]
	boards := prepareData(s[1:])
	for _, number := range strings.Split(numbers, ",") {
		for _, board := range boards {
			board.Play(string(number))
			win, val := board.HasWon()
			if win {
				return StrToInt(number) * val
			}
		}
	}
	return 0
}
