# Python – GlitchTip Integration

Dieses Beispiel zeigt, wie ein Python-Skript einen Fehler an GlitchTip senden kann.

## Installation

1. Python installieren (mindestens Version 3.x)
2. Abhängigkeit installieren:
```bash
pip install sentry-sdk
```

## Ausführen

```bash
python3 test_error.py
```

Das Skript löst absichtlich einen Fehler aus, der an GlitchTip gesendet wird (Projekt-ID #4).

## Konfiguration

In `test_error.py` befindet sich die DSN:
```python
sentry_sdk.init("https://6393d00179dc45218391f49fb45007b0@glitchtip-stg.puzzle.ch/4")
```
