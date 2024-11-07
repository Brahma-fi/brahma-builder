package utils

import (
	"encoding/json"
)

func Struct2Map(v interface{}) (map[string]interface{}, error) {
	r, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	out := make(map[string]interface{})
	if err = json.Unmarshal(r, &out); err != nil {
		return nil, err
	}

	return out, err
}

func ParseMemo(input string, target interface{}) error {
	return json.Unmarshal([]byte(input), target)
}
