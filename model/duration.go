package model

import (
	"strings"
	"strconv"
)
type Duration struct {
    int
}

func (dur *Duration) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
    d := strings.Replace(s,"s","",-1);
    if s == "null" {
       dur.int = 0;
       return
    }
    intVar, err := strconv.Atoi(d);
	dur.int = intVar
    return
}