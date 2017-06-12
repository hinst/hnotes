package huser

func AssertResult(e error) {
	if e != nil {
		panic(e)
	}
}