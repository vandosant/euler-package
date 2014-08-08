package euler

import (
	"strconv"
)

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

func PrimeFactorization(n int) []int {
	r := []int{}


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

func Factorization(n int) []int {
	r := []int{}
	var i int

	for i = 2; i < n; i++ {
		if n%i == 0 {
			r = append(r, i, (n/i))
			return r
		}
	}
	return r
}

func Prime(n int) bool {
	var i int

	for i = 2; i < n; i++ {

		if n%i == 0 && n != i {
			return false
		}
	}
	return true
}

func LargestPalindrome(n int) int {
	var r int

	ceilStr := ""
	for i := 0; i < n; i++ {
		ceilStr = ceilStr+"9"
	}
	ceil, err := strconv.Atoi(ceilStr)
	if err != nil {
		panic(err)
	}

	floorStr := "1"
	for i := 0; i < n-1; i++ {
		floorStr = floorStr+"0"
	}
	floor, err := strconv.Atoi(floorStr)
	if err != nil {
		panic(err)
	}

	for j := floor; j <= ceil; j++ {
		for k := floor; k <= ceil; k++ {
			m := j * k
			s := strconv.Itoa(m)

			runes := []rune(s)
			for q, w := 0, len(runes) - 1; q < w; q, w = q+1, w-1 {
				runes[q], runes[w] = runes[w], runes[q]
			}
			swapped := string(runes)

			if s == swapped && m > n {
				n = m
			}
		}
	}

	r = n
	return r
}
