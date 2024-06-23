package divide_n_conquer

import (
	"testing"

	"github.com/jzes/entendendo-algoritmos.git/toolBox"
)

func TestSum(t *testing.T) {
	expectedOutput := 15
	input := toolBox.MakeRange(0, 5)
	output := Sum(input)
	if output != expectedOutput {
		t.Errorf("Expected: %v, got %v", expectedOutput, output)
	}

}
