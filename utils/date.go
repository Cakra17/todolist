package utils

import (
	"time"

	"github.com/mergestat/timediff"
)


func GetTime(value string) string {
	layout := "2006-01-02 15:04:05.0000000 -0700 MST m=+0.000000000"
	parsedTime, _ := time.Parse(layout,value)
  res := timediff.TimeDiff(time.Now().Add(time.Until(parsedTime)))
	return res
}
