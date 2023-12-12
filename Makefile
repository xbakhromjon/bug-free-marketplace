-include .env
export

DOCKER_COMPOSE_FILE=docker-compose.yaml

arg = $(filter-out $@,$(MAKECMDGOALS))

# proto
.PHONY: proto-gen
proto-gen:
	./scripts/gen-proto.sh

# git submodule init 	
.PHONY: pull-proto
pull-proto:
	git submodule update --init --recursive

# go generate	
.PHONY: go-gen
go-gen:
	go generate ./...

.PHONY: start
start:
	@echo "Start Containers"
	docker-compose -f ${DOCKER_COMPOSE_FILE} up -d ${DOCKER_SERVICES}
	sleep 2
	docker-compose -f ${DOCKER_COMPOSE_FILE} ps

.PHONY: stop
stop:
	@echo "Stop Containers"
	docker-compose -f ${DOCKER_COMPOSE_FILE} stop ${DOCKER_SERVICES}
	sleep 2
	docker-compose -f ${DOCKER_COMPOSE_FILE} ps

.PHONY: stop
rm: stop
	@echo "Remove Containers"
	docker-compose -f ${DOCKER_COMPOSE_FILE} rm -v -f ${DOCKER_SERVICES}

.PHONY: migration-up
migration-up:
	@echo "Migrations Up"
	sleep 2
	docker-compose run --rm migrate -path=migrations/ -database='postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' up

.PHONY: migration-generate
migration-generate:
	@echo "Generation migration file $(name)"
	sleep 2
	docker-compose run --rm migrate create -ext sql -dir ./migrations -seq $(name)

.PHONY: mod-download
mod-download:
	@echo "Go mod download"
	sleep 2
	docker-compose exec app go mod download

.PHONY: go-get
go-get:
	@echo "Go get ${arg}"
	sleep 2
	docker-compose exec app go get -d ${arg}

.PHONY: mod-tidy
mod-tidy:
	@echo "Go mod tidy"
	sleep 2
	docker-compose exec app go mod tidy
