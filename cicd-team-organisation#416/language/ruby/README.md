


# Ruby – GlitchTip Integration

Dieses Projekt zeigt, wie man mit dem `sentry-ruby` SDK Fehler an GlitchTip sendet.

## Installation

1. Stelle sicher, dass Ruby und Bundler installiert sind.
2. Füge in die `Gemfile` folgende Zeile ein:
```ruby
gem "sentry-ruby"
```
3. Installiere die Abhängigkeiten:
```bash
bundle install
```

## Ausführen

```bash
ruby test.rb
```

Dies löst absichtlich einen Fehler aus, der an GlitchTip Projekt #10 gesendet wird.

## Konfiguration

In `test.rb` befindet sich die DSN:
```ruby
Sentry.init do |config|
  config.dsn = 'https://69d0738116074e24a0f6a0727834737e@glitchtip-stg.puzzle.ch/10'
end
```