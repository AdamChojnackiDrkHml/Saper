package main

import (
	"Saper/game"
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	controlChannel := make(chan bool, 1)
	go game.Timer(controlChannel)

	time.Sleep(0 * time.Second)
	controlChannel <- true
	board := game.CreateBoard(5, 5, 4)
	board.PrintBoard()
	board.PrintPlayerBoard()

	reader := bufio.NewReader(os.Stdin)
	for {

		txt, _, _ := reader.ReadLine()
		cords := strings.Split(string(txt), " ")
		x, err := strconv.Atoi(cords[0])
		if err != nil {
			break
		}
		y, err := strconv.Atoi(cords[1])
		if err != nil {
			break
		}

		board.CheckField(x, y)
		board.PrintPlayerBoard()

	}
}
