package network

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/getsentry/sentry-go"
	"github.com/valyala/fasthttp"
)

func ConvertRequest(ctx *fasthttp.RequestCtx) *http.Request {
	defer func() {
		if err := recover(); err != nil {
			sentry.Logger.Printf("%v", err)
		}
	}()

	r := new(http.Request)

	r.Method = string(ctx.Method())
	uri := ctx.URI()
	// Ignore error.
	r.URL, _ = url.Parse(fmt.Sprintf("%s://%s%s", uri.Scheme(), uri.Host(), uri.Path()))

	// Headers
	r.Header = make(http.Header)
	r.Header.Add("Host", string(ctx.Host()))
	ctx.Request.Header.VisitAll(func(key, value []byte) {
		r.Header.Add(string(key), string(value))
	})
	r.Host = string(ctx.Host())

	// Cookies
	ctx.Request.Header.VisitAllCookie(func(key, value []byte) {
		r.AddCookie(&http.Cookie{Name: string(key), Value: string(value)})
	})

	// Env
	r.RemoteAddr = ctx.RemoteAddr().String()

	// QueryString
	r.URL.RawQuery = string(ctx.URI().QueryString())

	// Body
	r.Body = io.NopCloser(bytes.NewReader(ctx.Request.Body()))

	return r
}
