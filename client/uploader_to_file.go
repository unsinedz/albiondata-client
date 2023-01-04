package client

import (
	"os"
	"strconv"
	"strings"
)

type toFileUploader struct {
	filePathFormat string
	iteration      int64
}

// newHTTPUploader creates a new HTTP uploader
func newToFileUploader(filePath string) uploader {
	return &toFileUploader{
		filePathFormat: strings.Replace(filePath, "fs://", "", -1),
		iteration:      0,
	}
}

func (u *toFileUploader) sendToIngest(body []byte, topic string) {
	filePath := strings.Replace(u.filePathFormat, "{0}", strconv.FormatInt(u.iteration, 10), -1)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	defer file.Close()
	_, err = file.Write(body)
	if err == nil {
		u.iteration++
	}
}
