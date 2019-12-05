package m

import "encoding/json"

// Define the resource structure
type Resource struct {
	Timestamp int64  `json:"timestamp"`
	URL       string `json:"url"`
}

func (r Resource) ToBytes() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return b
}