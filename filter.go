package funktors

func Filter[V any](v []V, a func(index int, v1 V) bool) []V {
	var fAll []V
	for i, v1 := range v {
		if a(i, v1) {
			fAll = append(fAll, v1)
		}
	}
	return fAll
}
