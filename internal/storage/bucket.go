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

func IsExistBucket(baseDir, bucketName string) (bool, error) {
	path := filepath.Join(baseDir, "buckets.csv")
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return false, err
	}

	for _, record := range records {
		if record[0] == bucketName {
			return true, nil
		}
	}
	return false, nil
}

// func DeleteBuckets(baseDir string) ([]BucketMeta, error) {
// 	// if IsBucketEmpty(baseDir) {

// 	// }
// }
