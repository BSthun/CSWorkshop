package functions

import (
	"strconv"

	"github.com/getsentry/sentry-go"
	"github.com/go-errors/errors"

	"backend/utils/value"
)

func ExtractYear(year string) *int64 {
	defer func() {
		if r := recover(); r != nil {
			hub := sentry.CurrentHub().Clone()
			hub.Scope().SetExtra("year", year)
			hub.CaptureException(errors.New(r))
			return
		}
	}()

	if val, err := strconv.ParseInt(year[0:4], 10, 64); err != nil {
		hub := sentry.CurrentHub().Clone()
		hub.Scope().SetExtra("year", year)
		hub.CaptureException(err)
		return value.Ptr[int64](0)
	} else {
		return &val
	}
}
