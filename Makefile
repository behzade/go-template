.PHONY: *

COMPOSE = docker compose -f ./build/docker-compose.yaml
CLI = $(COMPOSE) exec app go run /app/cmd/cli

build:
	$(COMPOSE) build
up:
	$(COMPOSE) up -d
attach:
	$(COMPOSE) up app
down:
	$(COMPOSE) down
restart:
	$(COMPOSE) restart
logs:
	$(COMPOSE) logs app --tail=20 --follow
bash:
	$(COMPOSE) exec app bash
update_entity:
	$(CLI) update_entity
get_alters:
	$(CLI) get_alters
apply_alters:
	$(CLI) apply_alters
