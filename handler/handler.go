package handler

import "log"

// Handler package entrypoint
func Handler(r string, h string, p string, lp string) {

	log.Println("Opening connection for:", h)
	createSession(getBastion(), h, p, lp)
}
