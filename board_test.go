package board_test

import (
	"testing"

	"github.com/roylee0704/tic-tac-toe"
)

func TestBoardSize(t *testing.T) {
	b := board.New()
	l := len(b)

	if l != 9 {
		t.Errorf("want: %d, got: %d", 9, l)
	}
}

func TestMark(t *testing.T) {

	b := board.New()

	tests := []struct {
		pos    int
		symbol rune
		want   rune
	}{
		{0, 'x', 'x'},
		{1, 'x', 'x'},
		{8, 'x', 'x'},
		{3, 'x', 'x'},
	}

	for i, t := range tests {
		b.Mark(t.pos, t.symbol)

	}
}
