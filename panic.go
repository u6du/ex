package ex

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
