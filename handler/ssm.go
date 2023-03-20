package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// Create SSM Session Output: returns a StreamURL and Token to open a WebSocket connection (SSM)
func createSession(t string, h string, p string, lp string) *ssm.StartSessionOutput {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	svc := ssm.NewFromConfig(cfg)

	fmt.Printf("Starting SSM Session @ %s:%s on %s", h, p, t)
	params := &ssm.StartSessionInput{
		Target:       aws.String(t),
		DocumentName: aws.String("AWS-StartPortForwardingSessionToRemoteHost"),
		Parameters: map[string][]string{
			"host":            {h},
			"portNumber":      {p},
			"localPortNumber": {lp},
		},
	}

	o, err := svc.StartSession(context.TODO(), params)
	if err != nil {
		panic(err)
	}

	return o
}
