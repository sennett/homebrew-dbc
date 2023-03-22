package handler

import (
	"context"

	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/ktr0731/go-fuzzyfinder"
)

type DB struct {
	Cluster  string
	Instance string
}

// Return DB Type list of Database ClusterId's, InstanceId's and Endpoints.
func getEndpoints() []string {

	var endpoints []string

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Panic("configuration error: " + err.Error())
	}

	svc := rds.NewFromConfig(cfg)

	log.Println("Grabbing RDS Endpoints...")
	params := &rds.DescribeDBInstancesInput{}
	instance_list, err := svc.DescribeDBInstances(context.TODO(), params)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range instance_list.DBInstances {
		endpoints = append(endpoints, *i.Endpoint.Address)
	}

	return endpoints
}

func FuzzEndpoints() string {

	endpoints := getEndpoints()

	idx, err := fuzzyfinder.Find(endpoints, func(i int) string {
		return endpoints[i]
	})
	if err != nil {
		log.Panic(err.Error())
	}

	return endpoints[idx]
}
