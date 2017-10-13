package jsonFormater

import (
	"bytes"
	"encoding/json"
	"io"
)

func GetFromReader(r io.Reader) (string, error) {
	var input map[string]interface{}
	if err := json.NewDecoder(r).Decode(&input); err != nil {
		return "", err
	}
	b := bytes.NewBufferString("")
	e := json.NewEncoder(b)
	e.SetEscapeHTML(false)
	e.SetIndent("", "\t")
	err := e.Encode(input)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
