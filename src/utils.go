package main

// Check the current error
func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckErrorInt(e int) {
	if e != 0 {
		panic(e)
	}
}
