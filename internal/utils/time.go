package utils

import (
	"encoding/csv"
	"os"
	"path/filepath"
)

func LastModicationTime(bucket, baseDir string) {
	DeleteBucketFromCSV(bucket, baseDir)
}
func DeleteBucketFromCSV(baseDir, bucket string) error {
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
		} //0 - bucket name , 1 - creation time  , 2 creation time , 3 active
		if rec[0] == bucket {
			// newRecords = append(newRecords, rec[:2])
			// LastModificationTime := time.Now()
			// newRecords = append(newRecords, LastModificationTime)
			temp = append(temp, res[4])
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
