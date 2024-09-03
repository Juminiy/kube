package minio_api

import (
	"github.com/brianvoe/gofakeit/v7"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
)

func NewCred(id, secret string) *miniocred.Value {
	return &miniocred.Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
	}
}

func randAccessKeyID() string {
	return gofakeit.Password(
		true,
		true,
		true,
		false,
		false,
		AccessKeyIDMaxLen,
	)
}

func randSecretAccessKey() string {
	return gofakeit.Password(
		true,
		true,
		true,
		false,
		false,
		SecretAccessKeyMaxLen,
	)
}
