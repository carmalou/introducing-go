package math

func Average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func Min(xs []float64) float64 {
	total := xs[0]
	for _, v := range xs {
		if total > v {
			total = v
		}
	}

	return total
}

func Max(xs []float64) float64 {
	total := xs[0]
	for _, v := range xs {
		if total < v {
			total = v
		}
	}

	return total
}
