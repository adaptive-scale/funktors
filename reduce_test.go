package funktors

import (
	"log"
	"testing"
)

type Value struct {
	Value1 string
	Value2 string
}

func TestReduce(t *testing.T) {
	t1 := Reduce([]Value{
		{Value1: "1", Value2: "2"},
		{Value1: "2", Value2: "3"},
	}, func(a Value) map[string]Value {
		return map[string]Value{
			a.Value1: a,
		}
	})

	log.Println(t1)
}
