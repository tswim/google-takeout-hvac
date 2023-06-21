package model

import (
	"strconv"
	"strings"
)
type Duration int

func (dur *Duration) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
    d := strings.Replace(s,"s","",-1);
    if s == "null" {
       *dur = Duration(0);
       return
    }
    intVar, err := strconv.Atoi(d);
	*dur = Duration(intVar)
    return
}