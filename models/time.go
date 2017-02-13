package models

import (
	"strconv"
	"time"
)

type KongTime struct {
	time.Time
}

func (t *KongTime) UnmarshalJSON(buf []byte) error {
	ts, err := strconv.Atoi(string(buf))
	if err != nil {
		return err
	}
	tt := time.Unix(0, 1000000*int64(ts))

	t.Time = tt
	return nil
}
