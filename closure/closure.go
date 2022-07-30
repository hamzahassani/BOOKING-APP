package closure

func Createclosure(a int) func() int {
	inc := func() int {
		gh := "test"
		a = a + 1
		return a
	}

	return inc
}
