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
func GenerateToken(host string, port string, region string, user string) {
	var dbEndpoint string = fmt.Sprintf("%s:%s", host, port)

	if user == "" {
		fmt.Println("")
		log.Println("No User provided (-u) - Please enter User to authenticate as:")

		fmt.Println("")
		var userInput string
		fmt.Scanln(&userInput)

		user = userInput
		fmt.Println("")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic("configuration error: " + err.Error())
	}

	log.Println(fmt.Sprintf("Creating token for %s:%s::%s - %s", user, host, port, region))

	authenticationToken, err := auth.BuildAuthToken(
		context.TODO(), dbEndpoint, region, user, cfg.Credentials)
	if err != nil {
		log.Panic("failed to create authentication token: " + err.Error())
	}

	fmt.Println("")

	fmt.Println("Authentication Token: ")

	fmt.Println("")
	fmt.Println(authenticationToken)
	fmt.Println("")
}
