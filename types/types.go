package types

import "encoding/json"

type GCI struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

func (g GCI) Marshal() ([]byte, error) {
	bs, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	return bs, nil
}
