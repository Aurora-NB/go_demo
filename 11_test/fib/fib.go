package fib

var m = make(map[int]int)

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	if v, ok := m[n]; ok {
		return v
	}
	mn := fib(n-1) + fib(n-2)
	m[n] = mn
	return mn
}
