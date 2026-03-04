# Games-Webshop
Dieser Webshop bietet unterschiedliche Computerspiele an.

# API
Folgende APIs sind verfügbar:

- `/auth/login`: Hier kann man sich bei dem Webshop anmelden um eine Bestellung in `/checkout/placeorder` erstellen zu können.
- `/auth/logout`: Hier kann man sich vom Webshop abmelden (bisher wird nur eine Logout Message ausgegeben, es erfolgt kein richtiger Logout).
- `/products`: Hier werden alle Produkte aufgelistet, die im Webshop existieren.
- `/products/{id}`: Hier können Details eines Produkts mit der ID `{id}` angesehen werden.
- `/checkout/placeorder`: Hier kann eine Bestellung erstellt werden (bisher wird nur eine Message ausgegeben, dass eine Bestellung erfolgt ist).

Der Webshop besteht aus den Services `auth-service`, `checkout-service` und `product-service`
- `auth-service` übernimmt die Anmeldung und Abmeldung von Usern. Dieser Service generiert ein `jwt`, das im `checkout-service` benötigt wird.
- `checkout-service` übernimmt die Erstellung einer Bestellung von Produkten. Aktuell wird nur eine Bestätigungsnachricht zurückgegeben. Der User muss über den `auth-service` angemeldet werden um einen Checkout durchführen zu können.
- `product-service` zeigt die Produktliste aber auch einzelne Produkte an. 

Der Code jedes Services besteht jeweils aus den Packages `main`, `internal` und `pkg`.
- `main`: Hier befinden sich alle API-Endpunkten
- `internal`: Hier befinden sich die unterschiedlichen Handler der API-Endpunkte
- `pkg`: Hier befinden sich die Hilfsfunktionen für die Services

Man findet im Projektordner unter `api-requests/collections/webshop` eine "Bruno-Collection" mit API-Requests um schnell Anfragen an die API senden zu können. Bruno ist eine leichtgewichtige Open-Source Alternative zu Postman.

# Docker
Die Images der Services findet man im Docker Hub unter `crmsn/auth-service:latest`, `crmsn/checkout-service:latest` und `crmsn/product-service:latest`

Die Container startet man mit `docker compose up`

Der Webshop ist dann über den Port 8080 erreichbar.

# CI/CD
Im Ordner `.github/workflows` findet man die GitHub Actions die für CI/CD benutzt werden.
- **`go.yml`**: Baut und testet den Webshop
- **`release-please.yml`**: Erstellt bei größeren Änderungen automatisch Tags mit den Versionsnummern der Services und eine Zusammenfassung der Änderungen in einem PR
- **`publish.yml`**: Baut nach einem PR von release-please die Docker Images der Services und pusht sie in Docker Hub unter einem neuen Tag mit der aktuellen Versionsnummer

# Kubernetes
Im Ordner `kubernetes` findet man die Deployments und die Services, die man anwenden kann.