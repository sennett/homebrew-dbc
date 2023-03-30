package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// Assert AWS Credentials are configured
func AssertCredentials() {

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic("configuration error: " + err.Error())
	}

	if cfg.Credentials != nil {
		svc := sts.NewFromConfig(cfg)
		params := sts.GetCallerIdentityInput{}

		caller_id, err := svc.GetCallerIdentity(context.TODO(), &params)
		if err != nil {
			log.Fatal("Failed to call identity ", err.Error())
		}

		fmt.Println("")
		log.Println(fmt.Sprintf("Using %s @ %s", *caller_id.UserId, cfg.Region))
	} else {
		log.Fatal("Configuration Error. No AWS credentials configured: " + err.Error())
	}

}
