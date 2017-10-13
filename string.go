package jsonFormater

import (
	"encoding/json"
	"io"
)

func getFromReader(r io.Reader) (string, error) {
	var input map[string]interface{}
	if err := json.NewDecoder(r).Decode(input); err != nil {
		return "", err
	}
	result, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		return "", err
	}
	return string(result), nil
}
