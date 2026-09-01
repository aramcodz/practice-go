package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	expected := 52
	result := Add(50, 2)

	t.Run("Addter test", func(t *testing.T) {
		if result != expected {
			t.Errorf("expected %d, got %d ", expected, result)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
