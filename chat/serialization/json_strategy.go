package serialization

import "encoding/json"

type JSONStrategy struct{}

func (j *JSONStrategy) Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}

func (j *JSONStrategy) Deserialize(data []byte, v any) error {
	return json.Unmarshal(data, v)
}
