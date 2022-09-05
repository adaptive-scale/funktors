package funktors

func Concat[V any](v1 []V, v2 []V) []V {
	return append(v1, v2...)
}
