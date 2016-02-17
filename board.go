package board

// Board describes a tic-tac-toe board
type Board struct {
	spaces [9]rune
}

// New returns empty 3 x 3 board.
func New() *Board {
	return &Board{}
}

func (b *Board) Spaces() []rune {
	s := []
	s = copy(s, b.spaces)
	return s
}

// Mark marks symbol on board
func (b *Board) Mark(pos int, symbol rune) {

	if pos > len(b.spaces) {
		pos = len(b.spaces) - 1
	} else if pos < 0 {
		pos = 0
	}

	b.spaces[pos] = symbol // should I throw error or fix them?
}

// func (b *Board) String() string {
// 	s := fmt.Sprintf("|X|X|O|")
// 	s += fmt.Sprintf("|X|O|X|")
// 	s += fmt.Sprintf("|_|X|O|")
//
// 	return s
// }
