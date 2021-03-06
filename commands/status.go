package commands

import (
	"fmt"
	"log"

	"github.com/Skarlso/go-furnace/config"
	"github.com/Skarlso/go-furnace/utils"
	"github.com/Yitsushi/go-commander"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/fatih/color"
)

// Status command.
type Status struct {
}

// Execute defines what this command does.
func (c *Status) Execute(opts *commander.CommandHelper) {
	stackname := config.STACKNAME
	sess := session.New(&aws.Config{Region: aws.String(config.REGION)})
	cfClient := cloudformation.New(sess, nil)
	client := CFClient{cfClient}
	stack := stackStatus(stackname, &client)
	info := color.New(color.FgWhite, color.Bold).SprintFunc()
	log.Println("Stack state is: ", info(stack.Stacks[0].GoString()))
}

func stackStatus(stackname string, cfClient *CFClient) *cloudformation.DescribeStacksOutput {
	descResp, err := cfClient.Client.DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackname)})
	utils.CheckError(err)
	fmt.Println()
	return descResp
}

// NewStatus Creates a new Status command.
func NewStatus(appName string) *commander.CommandWrapper {
	return &commander.CommandWrapper{
		Handler: &Status{},
		Help: &commander.CommandDescriptor{
			Name:             "status",
			ShortDescription: "Status of a stack.",
			LongDescription:  `Get detailed status of the stack.`,
			Arguments:        "",
			Examples:         []string{"status"},
		},
	}
}
