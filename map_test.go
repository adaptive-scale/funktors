package funktors

import (
	"log"
	"testing"
)

func TestMap(t *testing.T) {
	t1 := Map([]string{"asd", "asd"}, func(index int, a string) string {
		return a + "asd"
	})
	log.Println(t1)
}
