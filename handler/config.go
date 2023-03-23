package handler

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
)

// Assert AWS Credentials are configured
func AssertCredentials() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic("configuration error: " + err.Error())
	}

	if cfg.Credentials != nil {
		log.Println("Credentials found @ " + cfg.Region)
	} else {
		log.Fatal("Configuration Error. No AWS credentials configured: " + err.Error())
	}

}
