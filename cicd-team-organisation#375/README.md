# Error Tracking

* Daniel Tschan
* Christoph Raaflaub

## Stand

Dani und ich haben die [Resultate der Umfrage](https://docs.google.com/presentation/d/1BYrcFR8IneRTLrVPLDv7hnaNy-Ia-Nq-6jHMoS_TrA4/edit?usp=sharing) und die [Anforderungen von Ruby](https://codimd.puzzle.ch/TQQdfFVaQjaFbdXj05JeqQ) besprochen. 
Die [neuste Version](https://glitchtip.com/blog/2025-05-19-glitchtip-5-0-released) unterstützt 95% der Anforderungen unserer Benutzer. 

Danach können wir die finale Evaluation mit Ruby machen. 

* [ ] TODO: Mit Dani Tschan zusammen für die fehlenden 5% noch ein Issue bei ihnen erfassen (Alert Silencing).

## Ersatz

Reto Kupferschmid:
die Staging instanz https://glitchtip-stg.puzzle.ch/ ist upgedated, falls du da schon mal testen möchtest

## Anforderungen

* Clients für Java, JavaScript, Ruby, Python, Go, Rust, ...
  * OpenTelemetry, für Ruby vorstellbar
  * Metadaten:
    * Mandant/Tenant
    * Sprache
    * Browser, inkl. Betriebssystem, Auflösung, ...
    * Komponente
    * Betriebsumgebung (Dev, Int, Prod)
    * Requestparameter
    * Korrelations ID
    * Span und Backtrace
* Fehler automatisch gruppieren
  * Fingerprinting 
* Unterstützung für Frontend
  * Source Maps
* Mandantenfähigkeit
  * Für Ruby eher nice to have
  * Mehrere Projekte ja
  * Separation der Nutzergruppen nicht zwingend 
* Alerting
    * Mail
    * WebHooks, z.B. für Chat
    * Ruby nutzt momentan nur Mail
    * Notifikationen müssen nicht sofort, aber sollten mehrmals täglich erfolgen
* Web Frontend
  * Fehler manuell gruppieren
  * Gruppierung aufheben nice to have
  * Fehler browsen
  * Fehler suchen
    * Saved Searches nice to have 
  * Tagging und Filtering?
  * Möglichkeit Tickets zu kommentieren und/oder zu verlinken mit Issue Tracker, etc.
  * Settings zu Alerting
    * Issues für eine bestimmte Zeit oder Anzahl Alerts ignorieren
    * Opt out aus Alerting für einzelne Benutzer
    * Alerting Rules
  * Möglichkeit Issues zu resolven
    * Erkennung von Regression 
* Security
    * OAuth/OIDC
    * RBAC (nice to have)
