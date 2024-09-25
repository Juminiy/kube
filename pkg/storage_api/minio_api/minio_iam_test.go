package minio_api

import (
	"github.com/Juminiy/kube/pkg/log_api/stdlog"
	"github.com/Juminiy/kube/pkg/util"
	miniocred "github.com/minio/minio-go/v7/pkg/credentials"
	"strconv"
	"testing"
)

// +passed
func TestClient_CreateAccessKey(t *testing.T) {
	// create
	accessKey, err := testMinioClient.CreateAccessKey()
	stdlog.Info(accessKey)
	util.SilentHandleError("create access key error", err)

	// delete
	util.SilentHandleError("delete access key error",
		testMinioClient.DeleteAccessKey(accessKey.AccessKeyID),
	)
}

// +passed
func TestClient_DeleteAccessKey(t *testing.T) {

}

// +passed
func TestClient_SetAccessPolicy(t *testing.T) {
	util.SilentHandleError("create access policy error",
		testMinioClient.SetAccessPolicy(
			&PolicyConfig{
				BusinessUser: BusinessUser{
					Name: "chisatox0129",
					ID:   strconv.Itoa(33),
				},
				Cred: miniocred.Value{
					AccessKeyID: "uUDC29bGJj3v15K33rAmM1urgRk6c924eov0IrF6PZz3BnHj24",
				},
				BucketName: "s3fs-mount-bucket-chisato0129",
			},
		))
}

// +passed
func TestClient_DeleteAccessPolicy(t *testing.T) {
	util.SilentHandleError("delete access policy error",
		testMinioClient.DeleteAccessPolicy(
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
	userCred, err := testMinioClient.CreateAccessKey()
	util.SilentHandleError("create access key error", err)
	stdlog.Info(userCred.AccessKeyID)

	policyConfig := PolicyConfig{
		BusinessUser: businessUser,
		Cred:         userCred,
		BucketName:   "s3fs-mount-bucket-huamoyan666",
	}

	// create access policy
	util.SilentHandleError("create access policy",
		testMinioClient.SetAccessPolicy(&policyConfig))

}
