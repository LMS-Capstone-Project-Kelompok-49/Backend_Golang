package bucket

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func UploadFile(fileName string, filePath string, fileType string) (url string, err error) {
	ctx := context.Background()
	endpoint := "192.168.1.4:9000"
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

	if fileType == "video" {
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

		//up
		info, err := minioClient.FPutObject(ctx, vidBucket, fileName, filePath, minio.PutObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		log.Print(info)
		log.Printf("Successfully uploaded %s of size %d\n", fileName, info.Size)

		vidUrl := endpoint + vidBucket + fileName
		return vidUrl, nil
	}

	if fileType == "ppt" {
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

		//up
		info, err := minioClient.FPutObject(ctx, pptBucket, fileName, filePath, minio.PutObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		log.Print(info)
		log.Printf("Successfully uploaded %s of size %d\n", fileName, info.Size)

		pptUrl := endpoint + pptBucket + fileName

		return pptUrl, nil
	}
	return "", nil
}
