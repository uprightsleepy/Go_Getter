package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	_ "log"
)

func GetSecret(secretName string) (string, error) {
	svc := secretsmanager.New(session.Must(session.NewSession()), aws.NewConfig().WithRegion(AwsZone))
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return "", err
	}

	return *result.SecretString, nil
}
