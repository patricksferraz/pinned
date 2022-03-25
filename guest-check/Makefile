.PHONY: build

build:
	docker-compose build $(SERVICES)

.PHONY: status logs start stop clean

ps:
	docker-compose ps $(SERVICES)

logs:
	docker-compose logs -f $(SERVICES)

up:
	docker-compose up -d $(SERVICES)

start:
	docker-compose start $(SERVICES)

stop:
	docker-compose stop $(SERVICES)

down:stop
	docker-compose down -v --remove-orphans

attach:
	docker-compose exec $(SERVICE) bash

prune:
	docker system prune

.PHONY: test

test:
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes

.PHONY: gen gtest

gen:
	protoc \
	--go_out=proto/pb \
	--go_opt=paths=source_relative \
	--go-grpc_out=proto/pb \
	--go-grpc_opt=paths=source_relative \
	--proto_path=proto/protofiles \
	proto/protofiles/*.proto

gtest:
	go test -v -cover -coverprofile coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
