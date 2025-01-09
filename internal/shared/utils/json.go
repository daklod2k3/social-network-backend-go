package utils

import (
	"encoding/json"
	"strings"
)

// trim string before deserialize
func Deserialize(s string, rs any) error {
	s = trim(s)
	return json.Unmarshal([]byte(s), &rs)
}

func trim(s string) string {
	s = strings.TrimSpace(s)
	idx := strings.Index(s, "{")
	if idx > -1 {
		s = s[idx:]
	}
	idx = strings.LastIndex(s, "}")
	if idx > -1 {
		s = s[:idx+1]
	}
	return s
}
