package goaws

import (
	"log"

	"github.com/Skarlso/go_aws_mine/config"
	"github.com/Skarlso/go_aws_mine/errorhandler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// CreateEC2 testing AWS configuration.
func CreateEC2(ec2Config *config.EC2Config) {
	log.Println("Creating ec2 session.")
	sess := session.New(&aws.Config{Region: aws.String("eu-central-1")})
	ec2Client := ec2.New(sess, nil)
	runResult, err := ec2Client.RunInstances(&ec2.RunInstancesInput{
		ImageId:      &ec2Config.ImageID,
		DryRun:       &ec2Config.DryRun,
		MaxCount:     &ec2Config.MaxCount,
		MinCount:     &ec2Config.MinCount,
		InstanceType: &ec2Config.InstanceType,
		KeyName:      &ec2Config.KeyName,
		Monitoring:   &ec2.RunInstancesMonitoringEnabled{Enabled: &ec2Config.Monitoring.Enable},
	})
	log.Println("Instance created with id: ", *runResult.Instances[0].InstanceId)
	errorhandler.CheckError(err)
	ec2Id := aws.StringSlice([]string{*runResult.Instances[0].InstanceId})
	err = ec2Client.WaitUntilInstanceRunning(&ec2.DescribeInstancesInput{InstanceIds: ec2Id})
	if err != nil {
		errorhandler.CheckError(err)
	}
}