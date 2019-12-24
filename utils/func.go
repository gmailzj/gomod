package utils

import (
	"fmt"
	"time"
)

func Uniqid(prefix ...string) string {
	var p = ""
	for k, v := range prefix {
		if k == 1 {
			p = v
		}
	}
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", p, now.Unix(), now.UnixNano()%0x100000)
}
