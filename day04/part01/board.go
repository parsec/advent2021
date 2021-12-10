package main

type Board struct {
	rowCells    [CellsPerSide][CellsPerSide]int
	isUserBoard bool
	hasBingo    bool
}

// take input (either live input from a read function or input from an array) and populate a new board struct
func (b *Board) Populate() {
	//todo
}

// print the board and tell if it has bingo
func (b *Board) PrintBoard() {
	//todo
}

// sets number to 0 if it matches with the bingo call
func (b *Board) MarkNumberAsChecked(num int) {
	//todo
}

// tells if the board has bingo in a horizontal or diagonal line
func (b *Board) IsBingo() bool {
	//todo

	return true
}
