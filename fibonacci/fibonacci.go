package fibonacci

func Fibonacci() func() int {
	a := 1
	b := 1
	return func() int {
		tmp := a
		a = b
		b = a + tmp
		return tmp
	}
}
