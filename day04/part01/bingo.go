package main

const (
	// defines the number of rows and columns
	CellsPerSide = 5
)

type Bingo struct {
	Boards  []*Board
	IsBingo bool
}

// inits x number of boards and populates with data from input
func (b *Bingo) Init(boardCount int) {
	//todo
}

func main() {

}
