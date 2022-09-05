package funktors

func Every[V any](v []V, a func(index int, v1 V) bool) bool {
	for i, v1 := range v {
		if !a(i, v1) {
			return false
		}
	}
	return true
}
