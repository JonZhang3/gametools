package common

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

const layout = "2006-01-02 15:04:05"

type DateTime time.Time

func (t *DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(*t).Format(layout))
	return []byte(stamp), nil
}

func (t *DateTime) UnmarshalJSON(b []byte) error {
	ti, err := time.Parse(layout, string(b))
	if err != nil {
		return err
	}
	*t = DateTime(ti)
	return nil
}

func (t *DateTime) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*t = DateTime(nullTime.Time)
	return
}

func (t DateTime) Value() (driver.Value, error) {
	return time.Time(t).Format(layout), nil
}

func (DateTime) GormDataType() string {
	return "time"
}

func (t DateTime) String() string {
	return time.Time(t).Format(layout)
}
