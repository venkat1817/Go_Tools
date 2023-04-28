/*
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"

)

func main() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // Replace with your desired region
	})
	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
		return
	}

	// Create an RDS client
	svc := rds.New(sess)

	// Call the DescribeOrderableDBInstanceOptions API to get the available instance types
	input := &rds.DescribeOrderableDBInstanceOptionsInput{
		Engine: aws.String("mysql"), // Replace with the desired database engine
	}
	result, err := svc.DescribeOrderableDBInstanceOptions(input)
	if err != nil {
		fmt.Println("Failed to describe orderable DB instance options:", err)
		return
	}

	// Extract the instance types from the API response
	var instanceTypes []string
	for _, option := range result.OrderableDBInstanceOptions {
		if option.DBInstanceClass != nil {
			instanceTypes = append(instanceTypes, *option.DBInstanceClass)
		}
	}

	// Print the available instance types
	fmt.Println("Available RDS instance types:")
	for _, instanceType := range instanceTypes {
		fmt.Println(instanceType)
	}
}
*/

package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
)

func main() {
	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), /
	})
	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
		return
	}
	
	svc := rds.New(sess)
	DBInstanceIdentifier := "database-1"
	
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(DBInstanceIdentifier),
		''
	}
	result, err := svc.DescribeDBInstances(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(result.DBInstances) == 0 {
		fmt.Println(DBInstanceIdentifier)
		return
	}
	instanceType := aws.StringValue(result.DBInstances[0].DBInstanceClass)
	fmt.Println("RDS instance type:")
	fmt.Println(instanceType)
}
