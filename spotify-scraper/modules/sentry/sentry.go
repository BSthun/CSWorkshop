package sentry

import (
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"

	"backend/modules/config"
)

func Init() {
	hostname, err := os.Hostname()
	if err != nil {
		logrus.Fatal("UNABLE TO GET HOSTNAME: " + err.Error())
	}

	environment := "development"
	if config.C.Environment == 2 {
		environment = "production"
	}

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:                   config.C.SentryDsn,
		Debug:                 false,
		AttachStacktrace:      true,
		SampleRate:            0,
		EnableTracing:         true,
		TracesSampleRate:      config.C.SentryTracesSampleRate,
		TracesSampler:         nil,
		IgnoreErrors:          nil,
		SendDefaultPII:        false,
		BeforeSend:            nil,
		BeforeSendTransaction: nil,
		BeforeBreadcrumb:      nil,
		Integrations:          nil,
		DebugWriter:           nil,
		Transport:             nil,
		ServerName:            hostname,
		Release:               "",
		Dist:                  "",
		Environment:           environment,
		MaxBreadcrumbs:        0,
		MaxSpans:              0,
		HTTPClient:            nil,
		HTTPTransport:         nil,
		HTTPProxy:             "",
		HTTPSProxy:            "",
		CaCerts:               nil,
		MaxErrorDepth:         10,
	}); err != nil {
		logrus.Fatal("UNABLE TO INITIALIZE SENTRY: " + err.Error())
	}

	logrus.Info("INITIALIZED SENTRY")
}
