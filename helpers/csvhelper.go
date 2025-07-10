package helpers

import (
	"encoding/csv"
	"os"
)

func ReadCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	notes, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(notes) > 0 {
		return notes[1:], nil
	}
	return [][]string{}, nil
}

func WriteCSV(filepath string, header []string, notes [][]string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if header != nil {
		if err := writer.Write(header); err != nil {
			return err
		}
	}

	if err := writer.WriteAll(notes); err != nil {
		return err
	}
	return nil
}
