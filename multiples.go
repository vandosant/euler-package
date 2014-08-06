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

func PrimeFactorization(n int64) []int64 {
	r := []int64{}


	a := Factorization(n)
	r = append(r, a...)

	startLength := len(r)
	endLength := 99
	for startLength < endLength {
		startLength = len(r)
		for i, q := range r {
			if !Prime(q) {
				b := Factorization(q)

				r[i] = r[len(r)-1]
				r = r[:len(r)-1]

				r = append(r, b...)
			}

			if i == len(r)-1 {
				endLength = len(r)
			}
		}
	}

	return r
}

func Factorization(n int64) []int64 {
	r := []int64{}
	var i int64
	var j int64


	for i = 1; i < n; i++ {
		for j = n-1; j > 1; j-- {
			if i*j == n {
				r = append(r, i, j)
				return r
			}
		}
	}
	return r
}

func Prime(n int64) bool {
	var i int64

	for i = 2; i < n; i++ {

		if n%i == 0 && n != i {
			return false
		}
	}
	return true
}
