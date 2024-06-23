package divide_n_conquer

import (
	"testing"

	"github.com/jzes/entendendo-algoritmos.git/toolBox"
)

func TestBinaryS(t *testing.T) {
	list := toolBox.MakeRange(100, 110)

	bsearcher := NewBSearcher(list)

	position, err := bsearcher.Search(103)
	if err != nil {
		t.Error(err.Error())
	}

	if position != 3 {
		t.Errorf("want %d, got %d", 3, position)
	}
}

func TestBinarySNotFound(t *testing.T) {
	list := toolBox.MakeRange(100, 110)

	bsearcher := NewBSearcher(list)

	position, err := bsearcher.Search(111)
	if err == nil {
		t.Errorf("want error got %v", position)
	}
}
