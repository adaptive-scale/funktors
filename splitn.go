package funktors

func SplitN[V any](v []V, n int) [][]V {
	var tAll [][]V
	var t []V

	if len(v) == n {
		tAll = append(tAll, v)
		return tAll
	}

	for i, v1 := range v {
		if i%n == 0 {
			t = []V{}
		}
		t = append(t, v1)
		if i%n == n-1 {
			tAll = append(tAll, t)
		}
	}
	if len(t) > 0 {
		tAll = append(tAll, t)
	}
	return tAll
}
