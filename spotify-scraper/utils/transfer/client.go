package transfer

import (
	"net/http"
	"time"
)

func NewClient() *http.Client {
	return &http.Client{
		Timeout: 2 * time.Second,
	}
}
