package storage

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"time"
	"triple-s/internal/utils"
)

func ListObjects(baseDir, bucket, objectKey string) ([]ObjectMeta, error) {
	var objects []ObjectMeta

	path := filepath.Join(baseDir, bucket, objectKey)

	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return objects, nil
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
		objects = append(objects, ObjectMeta{
			Name:         r[0],
			CreationDate: r[1],
		})
	}

	return objects, nil
}

func IsExistObject(baseDir, bucketName, objectKey string) (bool, error) {
	path := filepath.Join(baseDir, bucketName, "objects.csv")
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
		if record[0] == objectKey {
			return true, nil
		}
	}
	return false, nil
}

func CreateObjectFile(baseDir, bucket, objetcKey string) error {
	path := filepath.Join(baseDir, bucket, objetcKey)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil

}

func DeleteObjectFromCSV(baseDir, bucket string) error {
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

	newRecords := make([][]string, 0, len(records))
	for _, rec := range records {
		if len(rec) == 0 {
			continue
		}
		if rec[0] == bucket {
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
func RemoveObjectFile(baseDir, bucket string) error {
	path := filepath.Join(baseDir, bucket)
	return os.RemoveAll(path)
}
func IsObjectEmpty(basedir, bucket string) (bool, error) {
	dir := filepath.Join(basedir, bucket)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, e := range entries {
		if e.Name() != "objects.csv" {
			return false, err
		}
	}
	return true, nil
}
func AddObject(baseDir, bucket, objectKey string, meta ObjectMeta) error {

	path := filepath.Join(baseDir, bucket, "objects.csv")
	creationTime := time.Now()
	cretion := creationTime.Format(time.RFC3339)
	err1 := utils.WriteDataToCsv([]any{meta.Name, meta.Size, meta.ContentType, cretion}, path)
	if err1 != nil {
		return err1
	}
	return nil
}
