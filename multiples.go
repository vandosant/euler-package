package euler

func SumMultiples(v int) int {
	x := 0
	for i := 0; i < v; i++ {
		if i % 3 == 0 {
			x += i
		} else if i % 5 == 0 {
			x += i
		}
	}

	return x
}
