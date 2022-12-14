package types

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
)

type Longitude float64

func (l *Longitude) Check() error {
	if *l > 180 || *l < -180 {
		return errors.InvalidParam{Param: []string{"lng"}}
	}

	return nil
}
