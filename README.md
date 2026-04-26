# Games-Webshop
Dieser Webshop bietet unterschiedliche Computerspiele an.

# API
Folgende APIs sind verfügbar:

- `/auth/login`: Hier kann man sich bei dem Webshop anmelden um eine Bestellung in `/checkout/placeorder` erstellen zu können. Dabei wird ein Token übergeben.
- `/auth/logout`: Hier kann man sich vom Webshop abmelden (bisher wird nur eine Logout Message ausgegeben, es erfolgt kein richtiger Logout).
- `/products`: Hier werden alle Produkte aufgelistet, die im Webshop existieren.
- `/products/{id}`: Hier können Details eines Produkts mit der ID `{id}` angesehen werden.
- `/checkout/placeorder`: Hier kann eine Bestellung erstellt werden (bisher wird nur eine Message ausgegeben, dass eine Bestellung erfolgt ist). Dafür wird der Token aus `/auth/login` benötigt.

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

Für die lokale Nutzung ist Docker, Docker-Compose und Go erforderlich
Die Container startet man mit `docker compose up`

Die Services sind unter folgenden Ports erreichbar:
- `auth-service`: 8080
- `product-service`: 8081
- `checkout-service`: 8082

# CI/CD
Im Ordner `.github/workflows` findet man die GitHub Actions die für CI/CD benutzt werden.
- **`go.yml`**: Baut und testet den Webshop
- **`release-please.yml`**: Erstellt bei größeren Änderungen automatisch Tags mit den Versionsnummern der Services und eine Zusammenfassung der Änderungen in einem PR
- **`publish.yml`**: Baut nach einem merge des PR von release-please die Docker Images der Services und pusht sie in Docker Hub unter einem neuen Tag mit der aktuellen Versionsnummer

# Kubernetes
Im Ordner `kubernetes` findet man die Deployments und die Services, die man anwenden kann.
Dafür benötigt man einen Kubernetes Cluster und kubectl.

# ArgoCD
Im Ordner `app` findet man die Apps die ArgoCD verwaltet

# OpenTofu
Im Ordner `tf` findet man die `.tf`-Files für das Anfordern von Ressourcen in Azure
Hierfür wird OpenTofu oder Terraform benötigt.
- `main.tf`: definiert die Ressourcen, die angefordert werden sollen
- `variables.tf` definiert Variablen, wie die Region und die Resource Group, um flexibel deployen zu können

# Makefile
Die Makefile hilft beim bauen und testen der Microservices des Webshops
Hierfür wird Go benötigt.
- `all`: Baut und testet die Services
- `build`: baut die Services
- `test`: testet die Services
- `clean`: löscht die Binaries die durch den `build`-Prozess erzeugt wurden

# just-File
Die just-File hilft beim deployen des Webshops und des LGTM-Stacks auf Azure.
Hierfür wird just, kubectl, OpenTofu oder Terraform und das azure-cli benötigt.
- `just deploy`: Fragt die Ressourcen aus den `.tf`-Files bei Azure an, ruft `just connect` auf, lädt ArgoCD auf dem Kubernetes Cluster runter und ruft `just sync` auf
- `just connect`: Verbindet sich mit Azure
- `just sync`: Die Apps `webshop`, `lgtm-stack` und `alloy-stack` werden auf dem Cluster angewendet und von ArgoCD verwaltet
- `just monitor`: Aktiviert Port-Forwarding für Grafana und öffnet ein Browser Fenster mit Grafana
- `just monitor-password`: Gibt das Passwort für Grafana aus
- `just destroy`: Löscht die angefragen Ressourcen in Azure
- `just status`: Gibt den Status aller Pods und PVCs von LGTM aus
- `just argocd-ui`: Vergibt ArgoCD eine externe IP
- `just argocd-password`: Gibt das Passwort für ArgoCD aus
- `just webshop`: Aktiviert Port-Forwarding für alle Microservices des Webshops
- `just products`: Aktiviert das Port-Forwarding für den `product-service`
- `just auth`: Aktiviert das Port-Forwarding für den `auth-service`
- `just checkout`: Aktiviert das Port-Forwarding für den `checkout-service`
- `just webshop-stop`: Stoppt das Port-Forwarding für alle Microservices des Webshops