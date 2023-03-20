package funktors

func GroupBy[T comparable, V any](v []V, a func(v1 V) map[T]V) map[T][]V {
	var mFinal = map[T][]V{}
	for _, val := range v {
		m := a(val)
		for k, v := range m {
			mFinal[k] = append(mFinal[k], v)
		}
	}
	return mFinal
}
