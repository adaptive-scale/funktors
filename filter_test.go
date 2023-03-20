package funktors

import (
	"log"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	input := []string{"test123", "123all", "testvalue"}
	values := Filter(input, func(index int, a string) bool {
		return strings.Contains(a, "test")
	})
	log.Println(values)
}
