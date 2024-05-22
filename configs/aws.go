package configs

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func ConnectAWS() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(EnvAWSRegion()),
	})
	if err != nil {
		log.Fatal(err)
	}
	return sess
}
