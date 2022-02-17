package game

import (
	"fmt"
	"time"
)

func Timer(control <-chan bool) {
	secondsCounter := 0
	for {
		select {

		case <-control:
			return
		default:
			fmt.Println(secondsCounter)
			secondsCounter++
			time.Sleep(1 * time.Second)
		}
	}
}
