package funktors

func Fill[V any](v []V, valueToBeFilled V, pos1 int, pos2 int) bool {
	for i := pos1; i <= pos2; i++ {
		v[i] = valueToBeFilled
	}
	return true
}
