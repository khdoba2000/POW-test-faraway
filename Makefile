
.PHONY: vendor
vendor: 
	go mod vendor
	
tidy:
	go mod tidy


client-build:
	go build -mod=vendor -trimpath -o bin/client cmd/client/main.go


server-build:
	go build -mod=vendor -trimpath -o bin/server cmd/server/main.go
