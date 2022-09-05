package funktors

import (
	"log"
	"testing"
)

func TestFilter(t *testing.T) {
	t1 := Filter([]string{"asd1", "asd"}, func(index int, a string) bool {
		return a == "asd"
	})
	log.Println(t1)
}
