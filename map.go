package funktors

func Map[T any, V any](v []V, a func(index int, v1 V) T) []T {
	var tAll []T
	for i, v1 := range v {
		tAll = append(tAll, a(i, v1))
	}
	return tAll
}
