include variables.mk

build-server:
	go build -o $(BUILD_DIR)/server.out $(SERVER_DIR)

build-client:
	go build -o $(BUILD_DIR)/client.out $(CLIENT_DIR)