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
		// если файла нет — это НЕ ошибка, просто пустой список
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
