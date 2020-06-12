package utils

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustDo(f func() error) func() {
	return func() {
		Must(f())
	}
}
