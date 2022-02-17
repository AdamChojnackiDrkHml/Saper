package main

import (
	"Saper/game"
	"time"
)

func main() {
	controlChannel := make(chan bool, 1)
	go game.Timer(controlChannel)

	time.Sleep(10 * time.Second)
}
