package test

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/stretchr/testify/require"
	"log"
	"net/url"
	"testing"
	"time"
)

func TestMinioClient(t *testing.T) {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "admin"
	secretAccessKey := "12345677"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup

	exist, err := minioClient.BucketExists(context.Background(), "testclient")
	require.NoError(t, err)
	require.True(t, exist)

	// Upload the zip file
	objectName := "file_name.mp4"
	filePath := "/Users/kobshobs/go/src/github.com/kobshobe/minio_test/test/jump_rope.mp4"
	contentType := "application/video"

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(context.Background(), "testclient", objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %+v\n", objectName, info)

	// Set request parameters for content-disposition.
	reqParams := make(url.Values)
	//reqParams	url.Values	额外的响应头，支持
	//response-expires，  到期时间
	//response-content-type，  响应内容的类型
	//response-cache-control，   缓存控制
	// response-content-disposition。  内容处理 例如，附件形式强制下载"attachment; filename=\"chat.mp4\""
	// reqParams.Set("response-content-type", "attachment; filename=\"chat.mp4\"")
	reqParams.Set("response-content-disposition", "	video/mpeg4")
	// Generates a presigned url which expires in a day.

	presignedURL, err := minioClient.PresignedGetObject(context.Background(), "testclient", objectName, time.Hour, reqParams)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully generated presigned URL", presignedURL)
}
