package encode

import "encoding/json"

type ReturnResult struct {
	Data   map[string]interface{} `json:"Data"`
	Errno  int                    `json:"Errno"`
	ErrMsg string                 `json:"ErrMsg"`
}

func Jsonencode(result interface{}) ([]byte, error) {
	return json.Marshal(result)
}
