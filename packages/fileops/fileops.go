package fileops

import (
	"encoding/json"
	"errors"
	"os"
)

func WriteToFile(fileName string, value []byte) error {
	err := os.WriteFile(fileName, value, 0644)
	if err != nil {
		return errors.New("you fucked up dude, something went wrong while writing to file")
	}
	return nil
}

func ReadFromFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("you fucked up dude, something went wrong while reading from file")
	}
	return data, nil
}

// JSONModifier modifies JSON data in a file.
// This function is currently a placeholder and needs to be implemented.
// It is intended to perform modifications on JSON data stored in a file.

func JSONEncoder(stream interface{}) ([]byte, error) {
	jsonData, err := json.MarshalIndent(&stream, "", "")
	if err != nil {
		return nil, errors.New("failed to marshal JSON")
	}
	return jsonData, nil
}

func JSONDecoder[T interface{}](data []byte) (*T, error) {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, errors.New("failed to marshal JSON")
	}
	return &result, nil
}
