package hn

func AssertResult(e error) {
	if e != nil {
		panic(e)
	}
}