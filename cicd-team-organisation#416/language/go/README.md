


# Go – GlitchTip Integration

Dieses Beispielprojekt zeigt, wie man einen Fehler in einer Go-Anwendung an GlitchTip sendet.

## Installation

1. Go installieren (falls nicht vorhanden):  
   [https://golang.org/doc/install](https://golang.org/doc/install)

2. Abhängigkeiten installieren:
```bash
go mod tidy
```

## Ausführen

```bash
go run main.go
```

Dies löst absichtlich einen Fehler aus, der an GlitchTip gesendet wird (Projekt-ID #5).

## Konfiguration

Die DSN ist in `main.go` direkt eingebettet:
```go
sentry.Init(sentry.ClientOptions{
  Dsn: "https://b20716182bd8449897d067ef2e27b6fa@glitchtip-stg.puzzle.ch/5",
})
```