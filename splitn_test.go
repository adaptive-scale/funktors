package funktors_test

import (
	"fmt"
	"github.com/adaptive-scale/funktors"
	"testing"
)

func TestSplitn(t *testing.T) {
	a := funktors.SplitN([]string{"asd1", "asd2", "asd3", "asd4", "asd5"}, 2)
	fmt.Println(a)
}
