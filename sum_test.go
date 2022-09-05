package funktors

import (
	"log"
	"testing"
)

func TestSum(t *testing.T) {
	t1 := Sum([]int{1, 2})
	log.Println(t1)
}
