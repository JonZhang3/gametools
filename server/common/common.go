package common

import (
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

//func (t *DateTime) Scan(value interface{}) error {
//	switch value.(type) {
//	case time.Time:
//		*t = DateTime(value.(time.Time))
//	case string:
//		ti, err := time.Parse(layout, value.(string))
//		if err != nil {
//			return err
//		}
//		*t = DateTime(ti)
//	default:
//		return errors.New("cannot convert to DateTime")
//	}
//	fmt.Printf("%s", time.Time(*t).Format(layout))
//	return nil
//}

func (t DateTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (*DateTime) GormDataType() string {
	return "datetime"
}
