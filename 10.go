package task

func Bucket(a []float64) map[int][]float64 {
	r := make(map[int][]float64)
	for _, v := range a {
		k := int(v) / 10 * 10
		r[k] = append(r[k], v)
	}
	return r
}
