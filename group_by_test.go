package funktors

import (
	"log"
	"testing"
)

func TestGroupBy(t *testing.T) {
	t1 := GroupBy([]Value{
		{Value1: "1", Value2: "2"},
		{Value1: "1", Value2: "3"},
	}, func(a Value) map[string]Value {
		return map[string]Value{
			a.Value1: a,
		}
	})

	log.Println(t1)
}
