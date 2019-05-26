package libs

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

type XTime time.Time

//func (t XTime) MarshalJSON() ([]byte, error) {
//	//do your serializing here
//	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
//	return []byte(formatted), nil
//}

func (t *XTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = XTime(now)
	return
}

func (t XTime) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t XTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
