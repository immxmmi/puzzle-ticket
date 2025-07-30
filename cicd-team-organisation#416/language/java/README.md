


# Java – GlitchTip Integration

Dieses Beispiel zeigt, wie du mit dem Sentry SDK für Java Fehler an GlitchTip sendest.

## Installation

1. Stelle sicher, dass Java und Maven installiert sind.
2. In deiner `pom.xml` befindet sich bereits die Abhängigkeit:
```xml
<dependency>
  <groupId>io.sentry</groupId>
  <artifactId>sentry</artifactId>
  <version>1.7.27</version>
</dependency>
```

## Ausführen

```bash
mvn compile
mvn exec:java -Dexec.mainClass=MyClass
```

## Konfiguration

Die DSN findest du in `MyClass.java`:
```java
Sentry.init("https://a1e3c23fc03643b0bffee804f5aa611b@glitchtip-stg.puzzle.ch/7");
```

Fehler werden mit `Sentry.capture(...)` an GlitchTip gesendet.