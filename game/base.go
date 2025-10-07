package game

type Game interface {
	Configure() error
	RunStep() error
	Summary() error
	Quit()
}
