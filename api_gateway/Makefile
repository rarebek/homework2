CURRENT_DIR=$(shell pwd)


build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto:
	protoc --go_out=genproto/user_service --go-grpc_out=genproto/user_service protos/user_service/user.proto


swag-gen:
	echo ${REGISTRY}
	swag init -g api/router.go -o api/docs