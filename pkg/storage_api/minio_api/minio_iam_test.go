package minio_api

import (
	"github.com/Juminiy/kube/pkg/util"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
	"strconv"
	"testing"
)

// +passed
func TestClient_CreateAccessKey(t *testing.T) {
	// create
	accessKey, err := _cli.CreateAccessKey()
	t.Log(accessKey)
	util.SilentFatalf("create access key error", err)

	// delete
	util.SilentFatalf("delete access key error",
		_cli.DeleteAccessKey(accessKey.AccessKeyID),
	)
}

// +passed
func TestClient_DeleteAccessKey(t *testing.T) {

}

// +passed
func TestClient_SetAccessPolicy(t *testing.T) {
	util.SilentFatalf("create access policy error",
		_cli.SetAccessPolicy(
			&PolicyConfig{
				BusinessUser: BusinessUser{
					Name: "juminiyx0129",
					ID:   strconv.Itoa(33),
				},
				Cred: miniocred.Value{
					AccessKeyID: "uUDC29bGJj3v15K33rAmM1urgRk6c924eov0IrF6PZz3BnHj24",
				},
				BucketName: "s3fs-mount-bucket-juminiy0129",
			},
		))
}

// +passed
func TestClient_DeleteAccessPolicy(t *testing.T) {
	util.SilentFatalf("delete access policy error",
		_cli.DeleteAccessPolicy(
			&PolicyConfig{
				BusinessUser: BusinessUser{
					Name: "huamoyan666",
					ID:   strconv.Itoa(55),
				},
				Cred: miniocred.Value{
					AccessKeyID: "EzS4Q57XE3oYYftOblPiyvrz8swewEd87YC23u4nN14KUF9s7E",
				},
				BucketName: "s3fs-mount-bucket-huamoyan666",
				PolicyName: "55-huamoyan666-s3fs-mount-bucket-huamoyan666-Cf7E5GxzB3k",
			}))
}

// FINISH: create an accessKey
// FINISH: create policy attach to accessKey
// +passed
func TestClient_AccessKeyWorkFlow(t *testing.T) {
	businessUser := BusinessUser{
		Name: "huamoyan666",
		ID:   strconv.Itoa(55),
	}

	//create access key
	userCred, err := _cli.CreateAccessKey()
	util.SilentFatalf("create access key error", err)
	t.Log(userCred.AccessKeyID)

	policyConfig := PolicyConfig{
		BusinessUser: businessUser,
		Cred:         userCred,
		BucketName:   "s3fs-mount-bucket-huamoyan666",
	}

	// create access policy
	util.SilentFatalf("create access policy",
		_cli.SetAccessPolicy(&policyConfig))

}
