package tools

func HandlePanicError(err error) {
	if err != nil {
		panic(err)
	}
}
