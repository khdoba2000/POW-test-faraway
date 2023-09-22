CURRENT_DIR=$(shell pwd)
.PHONY: vendor
vendor: 
	go mod vendor
	
tidy:
	go mod tidy


client-build:
	go build -mod=vendor -trimpath -o bin/client ${CURRENT_DIR}/cmd/client/main.go

server-build:
	go build -mod=vendor -trimpath -o bin/server ${CURRENT_DIR}/cmd/server/main.go


start-server:
	docker-compose up server


start-client:
	docker-compose up client
