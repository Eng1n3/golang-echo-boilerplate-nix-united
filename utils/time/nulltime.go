package utils

import (
	"errors"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	if value != nil {
		nt.Time, nt.Valid = time.Now(), false
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		nt.Time, nt.Valid = v, true
	case string:
		parsedTime, err := time.Parse("2006-01-02 15:04:05.999999999-07:00", v)
		if err != nil {
			return err
		}
		nt.Time, nt.Valid = parsedTime, true
	default:
		return errors.New("unsupported type")
	}
	return nil
}
