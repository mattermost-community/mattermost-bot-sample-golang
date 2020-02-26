package types

import (
	"bytes"
	"encoding/json"
	"io"
)

// JSON maps go type to a JSON struct
type JSON map[string]interface{}

// ToReader returns a io.Reader after encoding the map into JSON byte array
func (j JSON) ToReader() io.Reader {
	b := bytes.NewBuffer(nil)
	json.NewEncoder(b).Encode(j)
	return b
}

// Response stores the json
type Response struct {
	Data string `json:"data,omitempty"`
}
