package funktors

func ForEach[V any](v []V, a func(index int, v1 V)) {
	for i, val := range v {
		a(i, val)
	}
}
