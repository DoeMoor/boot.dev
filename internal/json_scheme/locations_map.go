package json_scheme

import (
	"encoding/json"
	"time"
)

type Locations struct { 
	// Count    int64       `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Results   []location `json:"results"`
	timeStamp time.Time
}

type location struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

func (loc *Locations) NewScheme( rawResp []byte) error {
	err := json.Unmarshal(rawResp,loc)
	if err != nil {
		return err
	}
	loc.timeStamp = time.Now()
	return nil
}
