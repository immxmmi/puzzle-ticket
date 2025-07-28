package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

const (
	defaultHost = "glitchtip:8000"
	defaultKey  = "<project_key>"
)

var (
	glitchtipHost       = getEnv("GLITCHTIP_HOST", defaultHost)
	glitchtipProjectKey = getEnv("GLITCHTIP_PROJECT_KEY", defaultKey)
	glitchtipDSN        = "http://" + glitchtipProjectKey + "@" + glitchtipHost + "/1"
)

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: glitchtipDSN,
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
		panic("💥 absichtlicher Fehler")
	})

	http.HandleFunc("/exception", func(w http.ResponseWriter, r *http.Request) {
		sentry.CaptureException(errors.New("⚠️ Exception ausgelöst über /exception"))
		fmt.Fprintln(w, "Exception gesendet.")
	})

	log.Println("🚀 Server läuft auf http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	sentry.CaptureException(errors.New("my error"))

	sentry.Flush(time.Second * 5)
}