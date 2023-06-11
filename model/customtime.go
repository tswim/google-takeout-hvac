package model
import (
	"time"
	"strings"
)

type CustomTime struct {
    time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
    s := strings.Trim(string(b), "\"")
    d := strings.Replace(s,"[UTC]","",-1);
    if s == "null" {
       ct.Time = time.Time{}
       return
    }
    ct.Time, err = time.Parse(time.RFC3339, d)
    return
}