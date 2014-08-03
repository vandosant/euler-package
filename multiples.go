package euler

func SumMultiples(v int) int {
	x := 0
	for i := 0; i < v; i++ {
		if i%3 == 0 {
			x += i
		} else if i%5 == 0 {
			x += i
		}
	}

	return x
}

func Fibonacci(v int) []int {
	x := 1
	y := 2
	r := []int{x, y}

	for i := 2; i < v; i++ {
		x = r[i-1]
		y = r[i-2]
		n := x + y
		r = append(r, n)
		if n > v {
			break
		}
	}

	return r
}

func SumEvenValues(s []int) int {
	r := 0
	for i := range s {
		if s[i]%2 == 0 {
			r += s[i]
		}
	}

	return r
}
