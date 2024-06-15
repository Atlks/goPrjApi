package libx

import "encoding/json"

// jsonEncode encodes a Go data structure to a JSON string
func JsonEncode(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// jsonDecode decodes a JSON string to a Go data structure
func JsonDecode(jsonStr string, result interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), result)
	if err != nil {
		return err
	}
	return nil
}
