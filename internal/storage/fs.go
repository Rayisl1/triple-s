package storage

import (
	"os"
	"path/filepath"
)

func IsBucketEmpty(basedir, bucket string) (bool, error) {
	dir := filepath.Join(basedir, bucket)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, e := range entries {
		if e.Name() != "object.go" {
			return false, err
		}
	}
	return true, nil
}
func RemoveAllFromBuckets(basedir, bucket string) error {
	return os.RemoveAll(filepath.Join(basedir, bucket))
}
