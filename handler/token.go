package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	_ "github.com/go-sql-driver/mysql"
)

func handleToken(host string, port string, region string, user string) string {
	var dbEndpoint string = fmt.Sprintf("%s:%s", host, port)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	authenticationToken, err := auth.BuildAuthToken(
		context.TODO(), dbEndpoint, region, user, cfg.Credentials)
	if err != nil {
		panic("failed to create authentication token: " + err.Error())
	}

	return authenticationToken
}
