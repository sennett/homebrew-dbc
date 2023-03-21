package handler

// Handler package entrypoint
func Handler(r string, h string, p string, lp string) {
	createSession(getBastion(), h, p, lp)
}
