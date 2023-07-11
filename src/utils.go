package app

// Check the current error
func Check(e error) {
	if e != nil {
		log.fatal(e)
	}
}
