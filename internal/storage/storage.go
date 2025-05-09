package storage

import (
	"encoding/json"
	"os"
)

func WriteResultsToFile(results map[string]bool) error {
	if err := os.MkdirAll("result", os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create("result/results.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(results)
}
