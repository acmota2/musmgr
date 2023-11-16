package model

import (
	"strings"
	"time"
)

type DayOfYear time.Time

func (t *DayOfYear) UnmarshalJSON(b []byte) error {
	initVal := strings.Trim(string(b), `"`)
	if initVal == "" || initVal == "null" {
		return nil
	}

	date, err := time.Parse("2006-01-02", initVal)
	if err != nil {
		return err
	}

	*t = DayOfYear(date)
	return nil
}

func (t DayOfYear) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02") + `"`), nil
}
