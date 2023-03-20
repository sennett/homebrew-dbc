package handler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// Return Bastion Instance by finding by `tag:service = bastion`
func getBastion(region string) string {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	svc := ec2.NewFromConfig(cfg)

	params := &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("tag:service"),
				Values: []string{"bastion"},
			},
		},
	}

	instance_list, err := svc.DescribeInstances(context.TODO(), params)
	if err != nil {
		panic(err)
	}

	return *instance_list.Reservations[0].Instances[0].InstanceId
}
