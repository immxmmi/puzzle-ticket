


# JavaScript – GlitchTip Integration

Dieses Projekt zeigt, wie man mit dem Sentry Browser SDK Fehler an GlitchTip meldet.

## Nutzung (CDN-Variante)

1. Öffne die Datei `index.html` im Browser.
2. Die Datei lädt automatisch das Sentry SDK via CDN und löst einen Testfehler aus.

## GlitchTip-Konfiguration

In `index.html` wird Sentry so initialisiert:
```js
Sentry.init({
  dsn: "https://473459e27d214bc39d7ac0fbc784d9e1@glitchtip-stg.puzzle.ch/8",
  tracesSampleRate: 0.01,
});
```

## Fehler-Test

Ein absichtlicher Fehler wird ausgelöst durch:
```js
myUndefinedFunction();
```

Öffne die Entwicklerkonsole → Fehler erscheint → wird an GlitchTip Projekt #8 gesendet.