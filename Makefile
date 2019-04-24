
IMAGE_NAME=content-service
IMAGE_TAG=1.0.0
DB_NAME=db-todo

all: build proto

build:
	docker build -t $(SERVICE_NAME) pkg/service/*/

proto:
	protoc -I api/proto/v1 api/proto/v1/content-service.proto --go_out=plugins=grpc:./pkg/api/v1