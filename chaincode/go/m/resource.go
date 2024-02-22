package m

import "encoding/json"

// Define the resource structure
type Resource struct {
	Timestamp int64  `json:"timestamp"`
	AssetID string `json:"assetID"`
	Owner string `json:"owner"`
	Balance int `json:"balance"`
	AppraisedValue int `json:"appraisedValue"`
}

func (r Resource) ToBytes() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return b
}

func NewResource(b []byte) (Resource, error) {
	r := Resource{}
	err := json.Unmarshal(b, &r)
	return r, err
}
