package handlers

import (
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_inst"
	"github.com/Juminiy/kube/pkg/image_api/harbor_api"
	"github.com/Juminiy/kube/pkg/storage_api/minio_api"
	"github.com/Juminiy/kube/pkg/storage_api/minio_api/minio_inst"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ImageUpload(c *gin.Context) {
	imageName := c.PostForm("image_name")
	imageTag := c.PostForm("image_tag")
	imageFileHdr, err := c.FormFile("image_file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	imageFile, err := imageFileHdr.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	defer util.SilentCloseIO("image file", imageFile)

	importResp, err := docker_inst.ImportImageV2((harbor_api.ArtifactURI{
		Project:    "hln7897",
		Repository: imageName,
		Tag:        imageTag,
	}).String(), imageFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, importResp)
}

func ImageDownloadURL(c *gin.Context) {
	fileName := c.PostForm("filename")
	if len(fileName) == 0 {
		c.JSON(http.StatusBadRequest, "filename not found")
		return
	}
	dlURL, err := minio_inst.TempGetObject(&minio_api.ObjectConfig{
		BucketName: minio_inst.GetPublicBucket(),
		ObjectPath: "",
		ObjectName: fileName,
	}, util.TimeSecond(600))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dlURL.String())
}
