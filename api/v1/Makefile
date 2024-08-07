all: service
MAIN_DIR := $(shell dirname -- $(shell dirname -- $(shell dirname -- $(shell pwd))))
USER_ID := $(shell id -u):$(shell id -g)
REGISTRY_NAME=servicehubregistry
IMG_NAME=service_hub_environment
IMG_DATE_TAG=20240329

# api.swagger.json file is automatically generated using npx buf generate command via the openapi plugin
# Set up goproxy variables in accordance to if goModuleNamePrefix is the user's repository or internal go.goms.io
service:
	set -e; \
	docker run --mount src=$(MAIN_DIR),target=/app/binded-data,type=bind $(REGISTRY_NAME).azurecr.io/$(IMG_NAME):$(IMG_DATE_TAG) /bin/bash -c \
	"cd binded-data/mygreeterv3/api/v1; \export GOPRIVATE=$(GOPRIVATE); export GOPROXY=$(GOPROXY); export GONOPROXY=$(GONOPROXY); export AKS_GOPROXY_TOKEN=$(AKS_GOPROXY_TOKEN); \
	cd proto; buf mod update; cd ..; buf generate; \
	swagger-codegen-cli generate -i api.swagger.json -l go -o restsdk -c swagger-config.json; \
	go mod tidy; \
	go get go.uber.org/mock/mockgen/model; \
	mockgen -package mock go.goms.io/aks/rp/mygreeterv3/api/v1 MyGreeterClient > mock/mock_client.go; \
	go mod tidy; \
	chown -R $(USER_ID) ../..; \
	cd /app; ./maintain_service populateMethodFiles --protoFilePath binded-data/mygreeterv3/api/v1/proto/api.proto --svcDirName binded-data/mygreeterv3 --deleteGarbageFiles=false;"; \