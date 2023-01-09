package loadgen

import (
	"fmt"
	"time"
)

type Gen struct {
	timeoutNS time.Duration
	lps       int
}

func Test(gen Gen) {
	gen.timeoutNS = 50 * time.Microsecond
	gen.lps = 1000
	var total64 = int64(gen.timeoutNS)/int64(1e9/gen.lps) + 1
	fmt.Println(total64)
}
