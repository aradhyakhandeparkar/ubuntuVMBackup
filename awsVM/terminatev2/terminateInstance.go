// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX - License - Identifier: Apache - 2.0
// snippet-start:[ec2.go-v2.CreateInstance]
package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// EC2TerminateInstanceAPI defines the interface for the TerminateInstances functions.
// We use this interface to test the functions using a mocked service.
type EC2TerminateInstanceAPI interface {
	TerminateInstances(ctx context.Context,
		params *ec2.TerminateInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.TerminateInstancesOutput, error)
}

// DeleteInstance terminates an Amazon Elastic Compute Cloud (Amazon EC2) instance.
// Inputs:
//
//	c is the context of the method call, which includes the AWS Region.
//	api is the interface that defines the method call.
//	input defines the input arguments to the service call.
//
// Output:
//
//	If success, a TerminateInstancesOutput object containing the result of the service call and nil.
//	Otherwise, nil and an error from the call to TerminateInstances.
func DeleteInstance(c context.Context, api EC2TerminateInstanceAPI, input *ec2.TerminateInstancesInput) (*ec2.TerminateInstancesOutput, error) {
	return api.TerminateInstances(c, input)
}

func DeleteInstanceCmd() {
	instanceId := flag.String("i", "", "The IDs of the instance to terminate")
	flag.Parse()

	if *instanceId == "" {
		fmt.Println("You must supply an instance ID (-i INSTANCE-ID")
		return
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ec2.NewFromConfig(cfg)

	instances := strings.Split(*instanceId, ",")

	input := &ec2.TerminateInstancesInput{
		InstanceIds: instances,
		DryRun:      new(bool),
	}

	result, err := DeleteInstance(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error terminating the instance:")
		fmt.Println(err)
		return
	}

	fmt.Println("Terminated instance with id: ", *result.TerminatingInstances[0].InstanceId)
}

// snippet-end:[ec2.go-v2.CreateInstance]
func main() {
	DeleteInstanceCmd()
}
