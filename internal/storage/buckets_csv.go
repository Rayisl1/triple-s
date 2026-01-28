package storage

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

func ListBuckets(baseDir string) ([]BucketMeta, error) {
	var buckets []BucketMeta

	path := filepath.Join(baseDir, "buckets.csv")

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return buckets, nil
		}
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, r := range records {
		if len(r) < 2 {
			continue
		}
		buckets = append(buckets, BucketMeta{
			Name:         r[0],
			CreationDate: r[1],
		})
	}

	return buckets, nil
}
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
