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

## ❓ Fragestellungen

### Allgemein & strategisch

- Was ist Apache Kafka genau – und wofür wird es eingesetzt?
- In welchen Szenarien ist Kafka besonders geeignet?
- Ist Kafka aktuell das "Tool to go" für Event-driven Architekturen?
- Welche Alternativen zu Kafka gibt es – und wann sind sie sinnvoller?
- Welche Vor- und Nachteile bringt Kafka im Vergleich zu anderen Message-Brokern?

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
