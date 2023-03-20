package handler

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func Handler(r string, h string, p string, lp string) {

	var session ssm.StartSessionOutput

	bastion := getBastion(r)
	session = handleSession(bastion, g, p, lp)

	handleSession(session.StreamURL, session.TokenValue)
}
