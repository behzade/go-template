.PHONY: *

COMPOSE = docker compose -f ./docker/docker-compose.yaml
CLI = $(COMPOSE) exec app go run /app
OAPI = $(COMPOSE) exec app oapi-codegen -generate
OAPI_TYPES = $(OAPI) types -o "./internal/controller/api-types.go" -package "controller" "./api/openapi/service.yaml" 
OAPI_SERVER = $(OAPI) server -o "./internal/controller/api.go" -package "controller" "./api/openapi/service.yaml" 

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
sync_entity:
	$(CLI) sync_entity
get_alters:
	$(CLI) get_alters
apply_alters:
	$(CLI) apply_alters
api:
	$(OAPI_TYPES) && $(OAPI_SERVER)
