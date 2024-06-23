package divide_n_conquer

import "fmt"

type BSearcher struct {
	itens []int
	start int
	end   int
}

func NewBSearcher(itens []int) *BSearcher {
	return &BSearcher{
		itens: itens,
		start: 0,
		end:   len(itens) - 1,
	}
}

// criar um tipo pra amarrar o start e o end do array
func (bs *BSearcher) Search(searched int) (int, error) {
	if bs.start > bs.end {
		return -1, fmt.Errorf("%v not found", searched)
	}

	midle := (bs.start + bs.end) / 2
	guess := bs.itens[midle]

	switch {
	case guess == searched:
		return midle, nil
	case guess > searched:
		bs.end = midle - 1
		return bs.Search(searched)
	}
	bs.start = midle + 1
	return bs.Search(searched)
}
