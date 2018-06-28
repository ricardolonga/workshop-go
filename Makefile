.DEFAULT_GOAL := help

.PHONY: all tools clean goget env env-ip test do-test env-stop test do-cover cover build image help

NAME    = workshop-go
VERSION = 1.0.0
GOTOOLS = \
	github.com/kardianos/govendor \
	golang.org/x/tools/cmd/cover

all: build image

tools: ## Instalar as ferramentas de cobertura e gestão de dependências
	go get -u -v $(GOTOOLS)

clean: ## Remover binário antigo
	-@rm -f $(NAME); \
	find vendor/* -maxdepth 0 -type d -exec rm -rf '{}' \;

goget: tools ## [tools] Baixar as dependências
	govendor sync -insecure +external

env: ## Subir ambiente necessário para os testes
	docker-compose up -d

env-ip: ## Visualizar o IP do container do MongoDB
	echo $$(docker inspect -f '{{.NetworkSettings.Networks.workshopgo_default.IPAddress}}' workshopgo_mongodb_1)

do-test: ## Executar os testes
	MONGO_URL=$$(docker inspect -f '{{.NetworkSettings.Networks.workshopgo_default.IPAddress}}' workshopgo_mongodb_1) go test $$(go list ./... | grep -v vendor)

env-stop: ## Finalizar ambiente necessário para os testes
	docker-compose kill
	docker-compose rm -f

test: env do-test env-stop ## [env do-test env-stop]

do-cover: ## Relatório de cobertura de testes
	./scripts/cover.sh

cover: env do-cover env-stop ## [env do-cover env-stop]

build: clean test ## [clean test] Construir o binário
	CGO_ENABLED=0 go build -v -a -installsuffix cgo -o $(NAME) ./cmd/server

image: ## Construir a imagem Docker
	docker build -t=$(NAME):$(VERSION) .

run-local: env
	MONGO_URL=$$(docker inspect -f '{{.NetworkSettings.Networks.workshopgo_default.IPAddress}}' workshopgo_mongodb_1) go run ./cmd/server/main.go

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
