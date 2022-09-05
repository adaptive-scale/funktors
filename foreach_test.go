package funktors

import (
	"fmt"
	"testing"
)

func TestForeach(t *testing.T) {
	ForEach([]string{"asd1", "asd"}, func(index int, a string) {
		fmt.Println("asd1" + a)
	})
}
