package utils

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"time"
)

func LastModicationTime(baseDir, bucket string) error {
	path := filepath.Join(baseDir, "buckets.csv")

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	var temp []string
	newRecords := make([][]string, 0, len(records))
	for _, rec := range records {
		if len(rec) == 0 {
			continue
		}
		if rec[0] == bucket {
			temp = append(temp, rec[0])
			temp = append(temp, rec[1])
			creationTime := time.Now()
			LastModificationTime := creationTime.Format(time.RFC3339)
			temp = append(temp, LastModificationTime)
			temp = append(temp, "active")
			newRecords = append(newRecords, temp)
			continue
		}
		newRecords = append(newRecords, rec)
	}
	tmpPath := path + ".tmp"
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(tmpFile)
	if err := writer.WriteAll(newRecords); err != nil {
		_ = tmpFile.Close()
		_ = os.Remove(tmpPath)
		return err
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		_ = tmpFile.Close()
		_ = os.Remove(tmpPath)
		return err
	}

	if err := tmpFile.Close(); err != nil {
		_ = os.Remove(tmpPath)
		return err
	}

	return os.Rename(tmpPath, path)
}
