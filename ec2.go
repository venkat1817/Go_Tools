package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := ec2.NewFromConfig(cfg)

	resp, err := client.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		ImageId:      types.ImageId("ami-0c55b159cbfafe1f0"),
		InstanceType: types.InstanceType("t2.micro"),
		MinCount:     1,
		MaxCount:     1,
	})
	if err != nil {
		log.Fatal(err)
	}

	instanceId := *resp.Instances[0].InstanceId
	fmt.Printf("Instance created with ID %s\n", instanceId)
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		resp, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
		if err != nil {
			log.Fatal(err)
		}

		var html string
		html += "<html><body><ul>"
		for _, reservation := range resp.Reservations {
			for _, instance := range reservation.Instances {
				html += fmt.Sprintf("<li>Instance ID: %s, State: %s</li>", *instance.InstanceId, instance.State.Name)
			}
		}
		html += "</ul></body></html>"

		fmt.Fprint(w, html)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
