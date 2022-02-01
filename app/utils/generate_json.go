package utils

import (
	"encoding/json"
	"errors"
)

func GenerateJson(param interface{}) ([]byte, error) {
	json, err := json.Marshal(param)

	if err != nil {
		return nil, errors.New("Error in convert json")
	}
	return json, nil
}
