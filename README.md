# Games-Webshop
Dieser Webshop bietet unterschiedliche Computerspiele an.

# API
Folgende APIs sind verfügbar:

- `/auth/login`: Hier kann man sich bei dem Webshop anmelden um eine Bestellung in `/checkout/placeorder` erstellen zu können.
- `/auth/logout`: Hier kann man sich vom Webshop abmelden (bisher wird nur eine Logout Message ausgegeben, es erfolgt kein richtiger Logout).
- `/products`: Hier werden alle Produkte aufgelistet, die im Webshop existieren.
- `/products/{id}`: Hier können Details eines Produkts mit der ID `{id}` angesehen werden.
- `/checkout/placeorder`: Hier kann eine Bestellung erstellt werden (bisher wird nur eine Message ausgegeben, dass eine Bestellung erfolgt ist).

Der Code besteht aus den Packages `main`, `internal` und `pkg`.
- `main`: Hier befindet sich der ausführbare Code mit allen API-Endpunkten
- `internal`: Hier befinden sich die unterschiedlichen Handler der API-Endpunkte
- `pkg`: Hier befinden sich die Hilfsfunktionen für den Webshop. Man findet Hilfsfunktionen für die Produkte in `products.go` und für den Token in `token.go`

Man findet im Projektordner unter `api-requests/collections/webshop` eine "Bruno-Collection" mit API-Requests um schnell Anfragen an die API senden zu können. Bruno ist eine leichtgewichtige Open-Source Alternative zu Postman

# Docker
Das Image des Webshops findet man unter `crmsn/devops-webshop-2026:latest`

Den Container startet man mit `docker run -p 8080:8080 devops-webshop-2026:latest`

Der Webshop ist über den Port 8080 erreichbar.