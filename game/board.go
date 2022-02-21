package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/kyokomi/emoji"
)

type boardS struct {
	bombsCounter  int
	width, height int
	score         int
	dataFields    [][]int
	fields        [][]FieldType
	playerFields  [][]string
}

func CreateBoard(width, height, bombsNum int) *boardS {
	newBoard := &boardS{
		bombsCounter: bombsNum,
		width:        width, height: height,
		score: 0, dataFields: make([][]int, height),
		fields:       make([][]FieldType, height),
		playerFields: make([][]string, height)}

	for i := range newBoard.dataFields {
		newBoard.dataFields[i] = make([]int, width)
		newBoard.fields[i] = make([]FieldType, width)
		newBoard.playerFields[i] = make([]string, width)

		for j := range newBoard.playerFields[i] {
			newBoard.playerFields[i][j] = "x"
		}
	}

	dataFields1D := make([]int, width*height)

	rand.Seed(time.Now().UnixNano())
	for bombsNum != 0 {
		leftFields := len(dataFields1D)
		for i := range dataFields1D {
			chances := float64(bombsNum) / float64(leftFields)
			if bombsNum == 0 {
				break
			}
			if chances > rand.Float64() {
				dataFields1D[i] = 9
				bombsNum--
			}
			leftFields--
		}
	}

	for i, f := range dataFields1D {
		newBoard.dataFields[i/height][i%width] = f
		newBoard.fields[i/height][i%width] = FieldType(f)
	}

	for i, row := range newBoard.dataFields {
		for j := range row {
			newBoard.dataFields[i][j] = newBoard.countBombs(i, j)
			newBoard.fields[i][j] = FieldType(newBoard.dataFields[i][j])
		}
	}

	return newBoard
}

func (board *boardS) countBombs(xCord, yCord int) int {
	if board.dataFields[xCord][yCord] == 9 {
		return 9
	}

	counter := 0

	neighbours := checkNeighborus(board.height, board.height, xCord, yCord)
	// fmt.Println(xCord, " ", yCord, ": ")
	for _, neighbour := range neighbours {
		// fmt.Print(xCord+neighbour[0], " ", yCord+neighbour[1], " ")
		if board.dataFields[xCord+neighbour[0]][yCord+neighbour[1]] == 9 {
			counter++
		}
	}
	// fmt.Println()
	return counter
}

func checkNeighborus(xLen, yLen, xCord, yCord int) (neighbours [][2]int) {

	if xCord != 0 && xCord != xLen-1 && yCord != 0 && yCord != yLen-1 {
		return [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}}
	}

	if xCord == 0 {
		if yCord == 0 {
			neighbours = [][2]int{{1, 0}, {1, 1}, {0, 1}}
		} else if yCord == yLen-1 {
			neighbours = [][2]int{{1, 0}, {1, -1}, {0, -1}}
		} else {
			neighbours = [][2]int{{1, 0}, {1, -1}, {0, -1}, {1, 1}, {0, 1}}
		}
	} else if xCord == xLen-1 {
		if yCord == 0 {
			neighbours = [][2]int{{-1, 0}, {-1, 1}, {0, 1}}
		} else if yCord == yLen-1 {
			neighbours = [][2]int{{-1, 0}, {-1, -1}, {0, -1}}
		} else {
			neighbours = [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {-1, 1}, {0, 1}}
		}
	} else {
		if yCord == 0 {
			neighbours = [][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}}
		} else {
			neighbours = [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}}
		}
	}

	return neighbours
}

func (board *boardS) PrintBoard() {
	for _, row := range board.dataFields {
		fmt.Println(row)
	}
}

func (board *boardS) PrintPlayerBoard() {

	for _, row := range board.playerFields {
		for _, ch := range row {
			if ch == "x" {
				emoji.Print(":green_square:")
				continue
			} else if ch == "O" {
				emoji.Print(":white_large_square:")
			} else if ch != ":bomb:" {
				emoji.Print(":keycap_" + ch + ": ")
			}
			///fmt.Print(ch + "\t")
		}
		fmt.Println()
	}
}

func (board *boardS) CheckField(xCord, yCord int) State {
	if xCord < 0 || xCord > board.height-1 || yCord < 0 || yCord > board.width-1 {
		return Invalid
	}

	if board.dataFields[xCord][yCord] == int(BOMB) {
		board.playerFields[xCord][yCord] = emoji.Sprint(":bomb:")
		return GameOver
	}
	if board.dataFields[xCord][yCord] == int(ZERO) {
		board.revealEmpty(xCord, yCord)
		return Valid
	}
	board.playerFields[xCord][yCord] = strconv.Itoa(board.dataFields[xCord][yCord])
	return Valid
}

func (board *boardS) CheckAllBombs() {
	for i, row := range board.dataFields {
		for j := range row {
			if board.dataFields[i][j] == int(BOMB) {
				board.playerFields[i][j] = emoji.Sprint(":bomb:")
			}
		}
	}
}

func (board *boardS) revealEmpty(xCord, yCord int) {

	if board.dataFields[xCord][yCord] != 9 && board.dataFields[xCord][yCord] != 0 {
		board.playerFields[xCord][yCord] = strconv.Itoa(board.dataFields[xCord][yCord])
		return
	}
	if board.playerFields[xCord][yCord] == "O" {
		return
	}
	board.playerFields[xCord][yCord] = "O"
	neighbours := checkNeighborus(board.height, board.width, xCord, yCord)

	for _, neigh := range neighbours {
		if board.dataFields[xCord+neigh[0]][yCord+neigh[1]] != int(BOMB) {
			board.revealEmpty(xCord+neigh[0], yCord+neigh[1])
		}
	}
}
