package divide_n_conquer

import "testing"

func TestBigger(t *testing.T) {
	expectedOutput := 5
	input := []int{2, 1, 4, 5, 2, 1}
	output := Bigger(input)
	if output != expectedOutput {
		t.Errorf("Expected: %v, got %v", expectedOutput, output)
	}

}
