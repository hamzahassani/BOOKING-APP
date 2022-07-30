package closure

func Createclosure(a int) func() int {
	inc := func() int {
		a = a + 1
		return a
	}

	return inc
}
