package main

import (
	"errors"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://2571f8176e854e3a98f33d14cb63cc5d@glitchtip-stg.puzzle.ch/6",
	})

	sentry.CaptureException(errors.New("my error test from Go"))

	sentry.Flush(time.Second * 5)
}
