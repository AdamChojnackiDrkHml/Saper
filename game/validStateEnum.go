package game

type State bool

const (
	Valid State = true
	Invalid  = false
	GameOver = false
)