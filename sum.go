package funktors

func Sum[V int64 | int32 | int | float64 | float32](v []V) V {

	var total V

	for _, val := range v {
		total += val
	}

	return total
}
