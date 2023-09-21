.PHONY: *

COMPOSE = docker compose -f ./build/docker-compose.yml --env-file ./.env

build:
	$(COMPOSE) build
up:
	$(COMPOSE) up -d
attach:
	$(COMPOSE) up
down:
	$(COMPOSE) down
restart:
	$(COMPOSE) restart
logs:
	$(COMPOSE) logs --tail=100
