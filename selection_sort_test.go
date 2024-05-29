package main

import (
	"reflect"
	"testing"
)

func TestLower(t *testing.T) {
	inputArr := []int{3, 4, 0, -1, 6, 8, -3}
	expOut := 6
	out := lower(inputArr)
	if out != expOut {
		t.Errorf("expected %d, got %d", expOut, out)
	}
}

func TestSelectSort(t *testing.T) {
	inputArr := []int{3, 4, 0, -1, 6, 8, -3}
	expOut := []int{-3, -1, 0, 3, 4, 6, 8}
	out := selectSort(inputArr)
	if !reflect.DeepEqual(expOut, out) {
		t.Errorf("expected %d, got %d", expOut, out)
	}
}
