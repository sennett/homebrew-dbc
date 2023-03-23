package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/aws/aws-sdk-go-v2/config"
)

// Create SSM Session Output: returns a StreamURL and Token to open a WebSocket connection (SSM)
func createSession(t string, h string, p string, lp string) {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic("configuration error: " + err.Error())
	}

	args := []string{
		"ssm",
		"start-session",
		"--region",
		cfg.Region,
		"--target",
		t,
		"--document-name",
		"AWS-StartPortForwardingSessionToRemoteHost",
		"--parameters",
		fmt.Sprintf("host=%s,portNumber=%s,localPortNumber=%s", h, p, lp),
	}

	command := exec.Command("aws", args...)

	stderr, err := command.StderrPipe()
	if err != nil {
		log.Panic(err.Error())
	}

	log.Printf("Connection Open at localhost:%s", lp)
	log.Println(command.Output())

	slurp, _ := ioutil.ReadAll(stderr)
	log.Printf("%s\n", slurp)
}
