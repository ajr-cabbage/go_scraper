package main

import (
	"encoding/json"
	"maps"
	"os"
	"slices"
)

func writeJSONReport(pages map[string]PageData, filename string) error {
	keys := slices.Sorted(maps.Keys(pages))
	sortedPages := []PageData{}
	for _, key := range keys {
		sortedPages = append(sortedPages, pages[key])
	}
	jsonPages, err := json.MarshalIndent(sortedPages, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, []byte(jsonPages), 0666)
	if err != nil {
		return err
	}

	return nil
}
