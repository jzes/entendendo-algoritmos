package main

import (
	"fmt"
)

func binaryS(itens []int, searched int) (int, error) {
	low := 0
	high := len(itens) - 1

	for low <= high {
		midle := (low + high) / 2
		guess := itens[midle]

		switch {
		case guess > searched:
			//se o chute é maior que o buscado, abaixa
			high = midle - 1
		case guess < searched:
			//se o chute é menor que o buscado, aumenta
			low = midle + 1
		case guess == searched:
			return midle, nil
		}
	}
	return 0, fmt.Errorf("[%v] not found", searched)
}
