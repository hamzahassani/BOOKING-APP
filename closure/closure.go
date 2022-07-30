package closure

func Createclosure(a int) func() int {
	inc := func() int {
		c := 12
		a = a + 1
		return a
	}

	return inc
}
