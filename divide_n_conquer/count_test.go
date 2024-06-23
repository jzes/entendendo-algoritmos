package divide_n_conquer

import "testing"

func TestCount(t *testing.T) {
	expectedOutput := 5
	input := []int{0, 0, 0, 0, 0}
	output := Count(input)
	if output != expectedOutput {
		t.Errorf("Expected: %v, got %v", expectedOutput, output)
	}
}
