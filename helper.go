package toml

import (
	"bytes"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := NewEncoder(&buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func MarshalString(v interface{}) (string, error) {
	b, err := Marshal(v)
	if err != nil {
		
		return "", err
	}
	return string(b), nil
}

func UnmarshalFile(path string, v interface{}) error {
	_, err := DecodeFile(path, v)
	return err
}
