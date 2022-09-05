package funktors

func FindAny[V any](v []V, a func(index int, v1 V) bool) V {
	for i, v1 := range v {
		if a(i, v1) {
			return v1
		}
	}
	return nil
}

func FindAll[V any](v []V, a func(index int, v1 V) bool) []V {
	return Filter(v, a)
}
