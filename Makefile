.PHONY: *

COMPOSE = docker compose -f ./build/docker-compose.yml

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
bash:
	$(COMPOSE) exec app bash
