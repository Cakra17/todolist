package utils

import "strconv"

func GenerateId(length int) string {
	if length == 0 {
    return "1"
  } else {
    return strconv.Itoa(length)
  }
}
