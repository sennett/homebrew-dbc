package handler

import (
	"context"
	"fmt"

	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	_ "github.com/go-sql-driver/mysql"
)

// Generate RDS IAM Authentication Token
func GenerateToken(host string, port string, region string, user string) string {
	var dbEndpoint string = fmt.Sprintf("%s:%s", host, port)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic("configuration error: " + err.Error())
	}

	authenticationToken, err := auth.BuildAuthToken(
		context.TODO(), dbEndpoint, region, user, cfg.Credentials)
	if err != nil {
		log.Panic("failed to create authentication token: " + err.Error())
	}

	return authenticationToken
}
