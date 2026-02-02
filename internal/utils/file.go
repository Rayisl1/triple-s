package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func WriteDataToCsv(data []any, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	row := make([]string, len(data))
	for i, v := range data {
		row[i] = fmt.Sprint(v)
	}

	return writer.Write(row)
}
