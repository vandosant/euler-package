package euler

func Multiples() int {
	x := 0
	for i := 0; i < 1000; i++ {
		if i % 3 == 0 {
			x += i
		} else if i % 5 == 0 {
			x += i
		}
	}

	return x
}
