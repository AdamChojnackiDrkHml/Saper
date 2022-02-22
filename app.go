package main

import (
	"Saper/game"
	"bufio"
	"os"
	"strconv"
	"time"
)

func main() {
	controlChannel := make(chan bool, 1)
	go game.Timer(controlChannel)

	N, err1 := strconv.Atoi(os.Args[1])
	M, err2 := strconv.Atoi(os.Args[2])
	B, err3 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil || err3 != nil {
		N, M, B = 5, 5, 4
	}
	time.Sleep(0 * time.Second)
	controlChannel <- true
	board := game.CreateBoard(N, M, B)
	board.PrintBoard()
	board.PrintPlayerBoard()

	reader := bufio.NewReader(os.Stdin)
	for {

		txt, _, _ := reader.ReadLine()
		state := board.InterpretCmd(string(txt))

		board.PrintPlayerBoard()

		if !state {
			break
		}

		if board.CheckWin() {
			break
		}
	}
}
