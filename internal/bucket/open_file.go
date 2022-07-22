package bucket

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func Open(file *multipart.FileHeader, fileType string, file_name string) (url string, err error) {
	fileSrc, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("error open file")
	}

	ext := filepath.Ext(file.Filename)

	file_name = fmt.Sprintf("%s%s", file_name, ext)

	filePath := filepath.Join("temp", file_name)

	fileDst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error")
	}

	if _, err := io.Copy(fileDst, fileSrc); err != nil {
		return "", fmt.Errorf("error")
	}

	fileUrl, err := UploadFile(file_name, filePath, fileType)
	if err != nil {
		return "", fmt.Errorf("error")
	} else {
		fileSrc.Close()
		fileDst.Close()

		err := os.Remove(filePath)
		if err != nil {
			return "", err
		}
	}
	return fileUrl, nil
}
