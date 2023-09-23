.PHONY: *

COMPOSE = docker compose -f ./build/docker-compose.yaml
ATLAS = $(COMPOSE) run --rm atlas --env dev

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
atlas_clean:
	$(ATLAS) schema clean
atlas_diff:
	$(ATLAS) schema apply --dry-run
