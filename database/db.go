package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	log "github.com/sirupsen/logrus"
)

// DB is the single database instance
var DB *dynamo.DB

func init() {
	log.Info("Initializing database")

	awsSession, err := session.NewSession()
	if err != nil {
		log.Fatalf("could not create a aws connection - %s", err)
	}
	DB = dynamo.NewFromIface(dynamodb.New(awsSession, &aws.Config{Region: aws.String("eu-west-1")}))
}
