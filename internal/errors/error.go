package errors

import "errors"

var (
	ErrInvalidBucketName   = errors.New("invalid bucket name")
	ErrBucketAlreadyExists = errors.New("bucket already exists")
	ErrBucketNotFound      = errors.New("bucket not found")
	ErrBucketNotEmpty      = errors.New("bucket not empty")

	ErrInvalidObjectKey = errors.New("invalid object key")
	ErrObjectNotFound   = errors.New("object not found")
)
