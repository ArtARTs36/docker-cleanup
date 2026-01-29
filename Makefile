.PHONY: up
up:
	docker-compose up

.PHONY: build
build:
	docker build . -t artarts36/docker-cleanup

.PHONY: lint
lint:
	golangci-lint run --fix
