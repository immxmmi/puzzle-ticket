package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("GLITCHTIP_DSN"),
		Environment: "development",
		Release:     "glitchtip-go@1.0.0",
	})
	if err != nil {
		log.Fatalf("Sentry Init fehlgeschlagen: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="de">
<head><meta charset="UTF-8"><title>GlitchTip-Test</title></head>
<body>
  <h1>GlitchTip Go Testseite</h1>
  <p>Nutze die Buttons, um Fehler zu erzeugen:</p>
  <button onclick="fetch('/error')">🔥 Panic Error</button>
  <button onclick="fetch('/exception')">⚠️ Capture Exception</button>
</body>
</html>`)
	})

	http.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		sentry.CaptureMessage("💥 Testfehler aus Go")
		fmt.Fprintln(w, "Fehler wurde ausgelöst.")
		go func() {
			panic("💥 absichtlicher Fehler")
		}()
	})

	http.HandleFunc("/exception", func(w http.ResponseWriter, r *http.Request) {
		sentry.CaptureException(errors.New("⚠️ Exception ausgelöst über /exception"))
		fmt.Fprintln(w, "Exception gesendet.")
	})

	log.Println("🚀 Server läuft auf http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}