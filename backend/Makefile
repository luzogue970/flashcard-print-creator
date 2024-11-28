# Variables
DOCKER_COMPOSE=docker-compose
API_CONTAINER_NAME=flashcards-api
DB_CONTAINER_NAME=flashcards-db

# Commandes principales
.PHONY: up down restart clean migrate

# Lancer les services
up:
	$(DOCKER_COMPOSE) up -d

# Arrêter les services
down:
	$(DOCKER_COMPOSE) down

# Relancer complètement (supprimer les données et les conteneurs)
restart: down clean up

# Supprimer le volume de données pour redémarrer à neuf
clean:
	$(DOCKER_COMPOSE) down -v
	docker volume prune -f

# Ajouter une commande de migration (exemple)
migrate:
# go run main.go
	go build -o myprogram.exe main.go
	.\myprogram.exe

# Raccourci pour relancer l'API et les services associés
run:
	$(MAKE) restart
	$(MAKE) migrate
