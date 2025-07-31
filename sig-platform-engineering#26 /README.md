# Kafka Knowhow-Aufbau & Admin-Fokus bei Puzzle

## ⚙️ Einordnung

Event-driven Design gewinnt in der modernen Softwarearchitektur zunehmend an Bedeutung – insbesondere im Kontext skalierbarer und entkoppelter Systeme. Technologien wie Apache Kafka spielen dabei eine zentrale Rolle als leistungsfähiger und bewährter Message-Broker.

Da unser Wissen im Bereich Kafka bei Puzzle etwas in den Hintergrund gerückt ist, haben David Simmen und ich (@mismail / Momo) entschieden, das Thema Kafka aktiv neu aufzugreifen und strukturiert weiterzuentwickeln. Die Verantwortlichkeiten teilen wir wie folgt auf:

- 🎯 **David** übernimmt den Fokus auf den Kafka **Developer-Bereich**
- 🛠️ **Ich (Momo)** kümmere mich – mit der Unterstützung von Falko (@fmacheleidt) – um den Kafka-**Administrationsbereich** und den gezielten Wissensaufbau

Unser Ziel ist es, dass Puzzle im Kafka-Umfeld wieder sichtbar wird – durch:

- internes technisches Know-how
- gezielte Labs und Demos
- dokumentierte Best Practices
- perspektivisch konkrete Angebote für Kundenprojekte

---

## 💬 Interne Kommunikation

Zur Koordination und zum Wissensaustausch wurde der Rocket.Chat-Kanal `#sig-kafka` erstellt.  
Alle interessierten Personen sind eingeladen, dort mitzudiskutieren, Fragen zu stellen oder Erfahrungen zu teilen.

---

## 📝 Summary

Der Spike soll klären, wie Kafka als Plattform administrativ effizient betrieben werden kann – mit Fokus auf Cluster-Setup, Sicherheit, Verwaltung und operativen Best Practices.  
Ziel ist ein strukturiertes Wissen, das zu Labs, Demos oder Angeboten ausgebaut werden kann.

---

## 🔍 Interner Abklärungsschritt

Bevor tief in die Technik eingestiegen wird, soll intern bei Puzzle geklärt werden:

- Wer hat bereits mit Kafka gearbeitet?
- Welche Kundenprojekte mit Kafka wurden bereits umgesetzt?
- Gibt es internes Wissen, Erfahrungsberichte oder Templates?
- Welche Lessons Learned oder Best Practices gibt es bereits?

Dieser Schritt dient dazu, vorhandenes Puzzle-Wissen zu nutzen und Doppelarbeit zu vermeiden.

---

## 🧭 Zielbild

- Aufbau eines fundierten Know-hows im Bereich Kafka-Administration
- Evaluierung und Dokumentation bewährter Tools und Vorgehensweisen
- Erstellung eines internen TechLabs oder Leitfadens
- 🔭 Langfristiges Ziel: Aufbau eines dedizierten Kafka-Angebots bei Puzzle (z. B. Beratung, Setup, Betrieb)

---



#### 🔍 Was ist Apache Kafka genau – und wofür wird es eingesetzt?

Apache Kafka ist eine skalierbare Event-Streaming-Plattform, die für hohe Durchsatzraten und niedrige Latenz bei der Datenverarbeitung entwickelt wurde. Sie ermöglicht:

- Hochperformante Datenübertragung zwischen Microservices oder Anwendungen
- Garantierte Nachrichtenreihenfolge
- Wiederholbare Verarbeitung durch Nachrichten-Replay
- Log-Komprimierung zur Optimierung des Speicherbedarfs
- Horizontale Skalierung und Datenreplikation zur Fehlertoleranz
- Lange Datenaufbewahrung für historische Analysen oder Replays




### Kafka ist besonders geeignet für Event-driven Architekturen, Event Sourcing, Monitoring, Log-Aggregation und Stream Processing in Echtzeit.

| Szenario                         | Beschreibung                                                                                   |
|----------------------------------|-----------------------------------------------------------------------------------------------|
| **Event-driven Architectures**   | Lose gekoppelte Systeme reagieren auf Ereignisse in Echtzeit – Kafka ist der zentrale Event-Bus. |
| **Event Sourcing**               | Applikationszustände werden über Ereignisse rekonstruierbar gespeichert.                        |
| **Log-Aggregation**              | Systeme (z. B. Microservices) senden Logs an Kafka; zentral gesammelt & weiterverarbeitet.     |
| **Monitoring & Metrics**         | Datenströme wie Metriken oder Traces können ingestet und verarbeitet werden.                   |
| **Stream Processing**            | Echtzeitverarbeitung großer Datenmengen (z. B. mit Kafka Streams, Flink).                      |
| **Data Ingestion Pipelines**     | Kafka als Puffer & Verteiler für heterogene Datenquellen (z. B. IoT, Web-Tracking, DB-Changefeeds). |
| **Messaging zwischen Systemen**  | Zuverlässige, skalierbare Kommunikation mit garantierter Reihenfolge.                          |
| **Audit / Compliance Logging**   | Unveränderliche Loghistorien für regulatorische Zwecke.                                        |


### 👉 Ist Apache Kafka weiterhin das "Tool to go" für Event-driven Architekturen – oder wann lohnen sich Alternativen wie Redpanda, NATS, RabbitMQ, Pulsar oder MQTT?

Apache Kafka ist weit verbreitet, aber nicht immer die beste Wahl:  
- **Redpanda**: Kafka-kompatibel, aber leichter, schneller und ohne JVM – ideal für latenzkritische oder ressourcenschwache Umgebungen  
- **Apache Pulsar**: Unterstützt Multi-Tenant-Architekturen und Trennung von Storage & Compute  
- **RabbitMQ**: Traditioneller Queue-basierter Broker mit starker Unterstützung für komplexes Routing  
- **NATS**: Extrem leichtgewichtig, ideal für einfache Messaging-Szenarien oder Edge-Umgebungen  
- **MQTT**: Für IoT/Embedded-Anwendungen mit minimalem Overhead

Die Wahl hängt stark vom Use Case ab – Kafka punktet bei großen Datenmengen, Replays und komplexem Event-Streaming.



### Welche Vor- und Nachteile bringt Kafka im Vergleich zu anderen Message-Brokern?

| Broker        | Persistente Speicherung | Replaying | Durchsatz | Latenz | Ressourcenverbrauch | Setup-Komplexität | Besonderheiten |
|---------------|--------------------------|-----------|-----------|--------|----------------------|--------------------|----------------|
| **Kafka**     | ✅ Ja                    | ✅ Ja     | 🔼 Hoch   | 🔽 Mittel | 🔼 Hoch              | 🔼 Hoch            | Reife Plattform mit großer Community |
| **Redpanda**  | ✅ Ja                    | ✅ Ja     | 🔼 Hoch   | 🔽 Sehr niedrig | 🔽 Gering          | 🔽 Mittel          | Kafka-kompatibel, kein JVM, C++-basiert |
| **Pulsar**    | ✅ Ja (tiered storage)   | ✅ Ja     | 🔼 Hoch   | 🔽 Mittel | 🔼 Mittel            | 🔼 Hoch            | Trennung von Compute und Storage |
| **RabbitMQ**  | 🟡 Eingeschränkt         | ❌ Nein   | 🔽 Mittel | 🔽 Niedrig | 🔽 Gering            | 🔽 Niedrig         | Komplexes Routing, AMQP |
| **NATS**      | ❌ Nein (nur mit JetStream) | 🟡 Teilweise | 🔼 Hoch   | 🔽 Sehr niedrig | 🔽 Sehr gering    | 🔽 Sehr gering     | Extrem leichtgewichtig |
| **MQTT**      | ❌ Nein                  | ❌ Nein   | 🔽 Gering | 🔽 Sehr niedrig | 🔽 Sehr gering      | 🔽 Sehr gering     | Speziell für IoT, minimales Protokoll |


### Wie Funkitoniert Kafka Primitiv


### Architektur & Betrieb

- Wie betreibt man Kafka zuverlässig in einer Cluster-Umgebung?
- Welche Deployment-Möglichkeiten gibt es (on-prem, Kubernetes, Cloud)?
- Wie lässt sich Kafka hochverfügbar und skalierbar betreiben?
- Welche Komponenten gehören zu einem vollständigen Kafka-Setup?
- Wie organisiert man ein Kafka-Cluster für mehrere Teams oder Anwendungen?

### Administration & Sicherheit

- Welche Best Practices gibt es für die Verwaltung von Topics und Benutzerrechten?
- Wie sichert man Kafka ab (z. B. Verschlüsselung, Authentifizierung, ACLs)?
- Wie geht man mit Quotas, Partitionierung und Datenaufbewahrung um?
- Welche Konventionen sind hilfreich für Namensgebung, Berechtigungen und Rollen?
- Wie behält man in einem wachsenden Kafka-Setup die Übersicht?

### Monitoring & Wartung

- Welche Tools eignen sich zur Überwachung von Kafka?
- Welche Metriken sind für Betrieb und Fehlersuche entscheidend?
- Wie erkennt und behebt man Probleme (z. B. Replikationsfehler, Lag)?
- Wie plant man ein Upgrade oder eine Migration (z. B. ZooKeeper → KRaft)?
- Wie führt man Backups durch und stellt Daten wieder her?

---

## 🎬 Showcase-Szenario


---

## 📘 Technische Dokumentation (Kurzfassung)

### 🔧 Komponentenübersicht

| Komponente      | Beschreibung                                      |
|------------------|---------------------------------------------------|
| Kafka-Cluster    | Verwaltet via Strimzi im Kubernetes-Cluster      |
| Producer         | Sendet Testnachrichten an definierte Kafka-Topics |
| Consumer         | Lauscht auf Topics zur Verifizierung              |
| Argo CD          | GitOps-basierte Verwaltung des Deployments        |
| Monitoring       | Grafana + Loki/Mimir für Logs & Metriken          |
| Authentifizierung| Optional via Keycloak (OAuth2 / OIDC)             |

### ⚙️ Technologiestack

- Kubernetes (K3s / Kind / GKE / OpenShift)
- Strimzi (Kafka mit KRaft oder ZooKeeper)
- GitOps mit Argo CD
- Grafana Stack (Mimir, Loki, Tempo)
- Optionale Security: TLS, SASL, SCRAM, ACLs

### 📂 Repositories & Tools

- Helm Charts: `charts/kafka-stack/`
- Beispiel-App: `examples/node-producer/`, `examples/go-consumer/`
- CI/CD via GitHub Actions oder GitLab CI

### 🗂️ Nächste Schritte

- ✅ Demo-Cluster aufsetzen (lokal oder remote)
- ✅ Producer/Consumer verifizieren
- ⏳ ACL-Konzept evaluieren & dokumentieren
- ⏳ Monitoring-Dashboard bauen
- ⏳ Doku & Lessons Learned strukturieren
