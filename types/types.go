package types

import (
	"encoding/json"
	"time"
)

type GCI struct {
	Name     string    `json:"name"`
	Data     []byte    `json:"data"`
	Previous time.Time `json:"previous_timestamp"`
	Current  time.Time `json:"current_timestamp"`
}

func (g GCI) Marshal() ([]byte, error) {
	bs, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	return bs, nil
}
