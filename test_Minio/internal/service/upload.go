package service

import (
	"mime/multipart"
)

type UploadService interface {
	UploadFile(file *multipart.FileHeader) (*string, error)
}
