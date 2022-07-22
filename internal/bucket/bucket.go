package bucket

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitMinio() *minio.Client {
	ctx := context.Background()
	// endpoint := "35.174.174.120:9000"
	// accessKeyID := "admin-capstone"
	// secretAccessKey := "SemogaLancar12345"

	endpoint := "192.168.127.217:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"

	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called video.
	vidBucket := "video"

	err = minioClient.MakeBucket(ctx, vidBucket, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, vidBucket)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", vidBucket)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", vidBucket)
	}

	//Make a new bucket called ppt.
	pptBucket := "ppt"

	err = minioClient.MakeBucket(ctx, pptBucket, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, pptBucket)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", pptBucket)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", pptBucket)
	}

	courseAvatar := "avatar"

	err = minioClient.MakeBucket(ctx, courseAvatar, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, courseAvatar)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", courseAvatar)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", courseAvatar)
	}

	mediaBucket := "media"

	err = minioClient.MakeBucket(ctx, mediaBucket, minio.MakeBucketOptions{})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, mediaBucket)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", mediaBucket)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", mediaBucket)
	}

	return minioClient
}

func UploadFile(fileName string, filePath string, fileType string) (url string, err error) {
	ctx := context.Background()

	minioClient := InitMinio()

	//up
	info, err := minioClient.FPutObject(ctx, fileType, fileName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	log.Print(info)
	log.Printf("Successfully uploaded %s of size %d\n", fileName, info.Size)

	fileUrl := fmt.Sprintf("%s/%s/%s", minioClient.EndpointURL(), fileType, fileName)
	return fileUrl, nil

}

func RemoveFile(fileName string, fileType string) error {
	log.Print("deleting ", fileName)
	ctx := context.Background()
	minioClient := InitMinio()

	err := minioClient.RemoveObject(ctx, fileType, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	log.Printf("Successfully deleted %s", fileName)
	return nil
}
