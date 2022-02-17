package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
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
		for i, _ := range dataFields1D {
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
		for j, _ := range row {
			newBoard.dataFields[i][j] = checkNeighborus(newBoard.dataFields, i, j)
			newBoard.fields[i][j] = FieldType(newBoard.dataFields[i][j])
		}
	}

	return newBoard
}

func checkNeighborus(fields [][]int, xCord, yCord int) int {
	if fields[xCord][yCord] == 9 {
		return 9
	}
	var neighbours [][2]int
	counter := 0
	if xCord != 0 && xCord != len(fields)-1 && yCord != 0 && yCord != len(fields[0])-1 {
		neighbours = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, -1}, {0, 1}}
	}

	if xCord == 0 {
		if yCord == 0 {
			neighbours = [][2]int{{1, 0}, {1, 1}, {0, 1}}
		} else if yCord == len(fields[0])-1 {
			neighbours = [][2]int{{1, 0}, {1, -1}, {0, -1}}
		} else {
			neighbours = [][2]int{{1, 0}, {1, -1}, {0, -1}, {1, 1}, {0, 1}}
		}
	} else if xCord == len(fields)-1 {
		if yCord == 0 {
			neighbours = [][2]int{{-1, 0}, {-1, 1}, {0, 1}}
		} else if yCord == len(fields[0])-1 {
			neighbours = [][2]int{{-1, 0}, {-1, -1}, {0, -1}}
		} else {
			neighbours = [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {-1, 1}, {0, 1}}
		}
	} else {
		if yCord == 0 {
			neighbours = [][2]int{{-1, 0}, {-1, 1}, {0, 1}, {-1, 1}, {-1, 0}}
		} else {
			neighbours = [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}}
		}
	}

	for _, neighbour := range neighbours {
		if fields[xCord+neighbour[0]][yCord+neighbour[1]] == 9 {
			counter++
		}
	}
	return counter
}

func (board *boardS) PrintBoard() {
	for _, row := range board.dataFields {
		fmt.Println(row)
	}
}

func (board *boardS) PrintPlayerBoard() {
	fmt.Print("T\t0\t1\t2\t3\t4\n")
	for i, row := range board.playerFields {
		fmt.Printf("%d\t", i)
		for _, ch := range row {
			fmt.Print("[" + ch + "]\t")
		}
		fmt.Println()
	}
}

func (board *boardS) CheckField(xCord, yCord int) bool {
	if xCord < 0 || xCord > board.height-1 || yCord < 0 || yCord > board.width-1 {
		return false
	}
	board.playerFields[xCord][yCord] = strconv.Itoa(board.dataFields[xCord][yCord])
	return true
}
