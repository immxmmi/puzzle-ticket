# Rust – GlitchTip Integration

Dieses Projekt zeigt, wie man mit dem Sentry SDK für Rust Fehler an GlitchTip sendet.

## Installation

1. Rust installieren (falls nicht vorhanden):  
   [https://www.rust-lang.org/tools/install](https://www.rust-lang.org/tools/install)

2. Projekt initialisieren (falls noch nicht geschehen):
```bash
cargo init
```

3. SDK hinzufügen:
```bash
cargo add sentry
```

## Ausführen

```bash
cargo run
```

Dies löst absichtlich einen Fehler aus, der an GlitchTip Projekt #9 gesendet wird.

## Konfiguration

In `main.rs` befindet sich die DSN:
```rust
let _guard = sentry::init("https://7c9cdfb30fe04f248197a6f684d279ce@glitchtip-stg.puzzle.ch/9");
```
