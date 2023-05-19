package text

import "time"

var BangkokTime *time.Location

func init() {
	BangkokTime, _ = time.LoadLocation("Asia/Bangkok")
}
