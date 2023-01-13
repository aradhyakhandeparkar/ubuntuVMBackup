package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"

	"fmt"
)

func TerminateInstance(client *ec2.EC2, instanceId string) error {
	_, err := client.TerminateInstances(&ec2.TerminateInstancesInput{
		InstanceIds: []*string{&instanceId},
	})

	return err
}

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}

	ec2Client := ec2.New(sess)

	instanceId := "i-067fedb079b00d57b"
	err = TerminateInstance(ec2Client, instanceId)
	if err != nil {
		fmt.Printf("Couldn't terimate instance: %v", err)
	}

	fmt.Println("Terminated instance with id: ", instanceId)
}
