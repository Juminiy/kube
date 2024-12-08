package handlers

import (
	"github.com/Juminiy/kube/cmd/deployof/sdk"
	"github.com/Juminiy/kube/pkg/image_api/docker_api"
	"github.com/Juminiy/kube/pkg/image_api/docker_api/docker_inst"
	"github.com/Juminiy/kube/pkg/storage_api/minio_api"
	"github.com/Juminiy/kube/pkg/storage_api/minio_api/minio_inst"
	"github.com/Juminiy/kube/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v3"
	"github.com/minio/minio-go/v7"
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

	importResp, err := docker_inst.ImportImageV3(
		sdk.GetImageAddr(imageName, imageTag), imageFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, importResp)
}

func ImageDownloadURL(c *gin.Context) {
	imageName := c.PostForm("image_name")
	imageTag := c.PostForm("image_tag")
	if len(imageName) == 0 || len(imageTag) == 0 {
		c.JSON(http.StatusBadRequest, "Form Param is error")
		return
	}

	var resp = struct {
		ExportInfo docker_api.ExportImageRespV2
		UploadInfo minio.UploadInfo
		FileURL    string
	}{}
	objCfg := minio_api.ObjectConfig{
		BucketName: minio_inst.GetPublicBucket(),
		ObjectPath: "",
		ObjectName: namedTarFile(imageName, imageTag),
	}
	found, reqErr := minio_inst.ObjectExists(&objCfg)
	if reqErr != nil {
		c.JSON(http.StatusInternalServerError, reqErr.Error())
		return
	}
	if !found {
		exportResp, err := docker_inst.ExportImageV2(sdk.GetImageAddr(imageName, imageTag))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		resp.ExportInfo = exportResp
		resp.UploadInfo, err = minio_inst.PutObject(&objCfg, exportResp.ImageFileReader)
		if err != nil {
			c.JSON(fiber.StatusInternalServerError, err.Error())
			return
		}
	}

	dlURL, err := minio_inst.TempGetObject(&objCfg, util.TimeSecond(600))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	resp.FileURL = dlURL.String()

	c.JSON(http.StatusOK, resp)
}

func namedTarFile(imageName, imageTag string) string {
	return imageName + "-" + imageTag + ".tar"
}
