MAKEFILE_THIS := $(lastword $(MAKEFILE_LIST))
SRC_PATH      := $(patsubst %/,%,$(dir $(abspath $(MAKEFILE_THIS))))
APP_FILE      := marble.json
APP_NAME      := marble_cc
BUILD_PATH    ?= /tmp/marble
FAB_PATH      ?= $(GOPATH)/src/github.com/hyperledger/fabric-samples
CC_DEPLOY     := $(FAB_PATH)/chaincode
SHIM_PATH     := $(SRC_PATH)/../../shim
PATCH_PATH    := $(SRC_PATH)/../../flogo-patch

CLIENT_FILE   := marble_client.json
CLIENT_PATH   ?= /tmp/marble_client
CLIENT_NAME   := marble_client

FE_CLIENT_FILE := marble_client_fe.json

FLOGO_VER     ?= v0.9.4
DT_REPO       ?= github.com/yxuco/hlf-contrib
FE_HOME       ?= $(HOME)/tibco/flogo28/2.8

.PHONY: all
all: create build

.PHONY: clean
clean:
	rm -Rf $(BUILD_PATH)

.PHONY: create
create: $(APP_FILE) clean
	mkdir -p $(BUILD_PATH)
	cp $(APP_FILE) $(BUILD_PATH)
	cd $(BUILD_PATH) && \
	flogo create --cv $(FLOGO_VER) -f $(APP_FILE) $(APP_NAME)
	rm $(BUILD_PATH)/$(APP_NAME)/src/main.go
	cp $(SHIM_PATH)/chaincode_shim.go $(BUILD_PATH)/$(APP_NAME)/src/main.go
	cp -Rf META-INF $(BUILD_PATH)/$(APP_NAME)/src
	cp $(PATCH_PATH)/codegen.sh $(BUILD_PATH)/$(APP_NAME)
	cd $(BUILD_PATH)/$(APP_NAME) && ./codegen.sh $(FE_HOME) && chmod +x src/gomodedit.sh

.PHONY: build
build: $(BUILD_PATH)/$(APP_NAME)
	cd $(BUILD_PATH)/$(APP_NAME)/src && \
	go mod edit -replace=github.com/project-flogo/flow@v0.9.4=github.com/yxuco/flow@v0.9.5 && \
	go mod edit -replace=github.com/project-flogo/core@v0.9.4=github.com/yxuco/core@v0.9.5

	-cd $(BUILD_PATH)/$(APP_NAME) && flogo build -e
	
	cd $(BUILD_PATH)/$(APP_NAME)/src && \
	go mod vendor && \
	env GOOS=linux GOARCH=amd64 go build -mod vendor -o ../$(APP_NAME)

.PHONY: deploy
deploy:
	rm -f $(CC_DEPLOY)/$(APP_NAME)_1.0.cds
	rm -Rf $(CC_DEPLOY)/$(APP_NAME)
	cp -Rf $(BUILD_PATH)/$(APP_NAME)/src $(CC_DEPLOY)/$(APP_NAME)
	cp fn-cli-init.sh $(FAB_PATH)/first-network/scripts
	cp fn-cli-test.sh $(FAB_PATH)/first-network/scripts

.PHONY: cli-init
cli-init:
	docker exec cli bash -c 'cd scripts; ./fn-cli-init.sh'

.PHONY: cli-test
cli-test:
	docker exec cli bash -c 'cd scripts; ./fn-cli-test.sh'

.PHONY: metadata
metadata:
	fabric-tools metadata -f $(APP_FILE) -m $(SRC_PATH)/contract-metadata/metadata.json -g $(SRC_PATH)/contract-metadata/metadata.gql -o $(SRC_PATH)/contract-metadata/override.json

.PHONY: clean-client
clean-client:
	rm -Rf $(CLIENT_PATH)

.PHONY: create-client
create-client: $(CLIENT_FILE) clean-client
	mkdir -p $(CLIENT_PATH)
	cp $(CLIENT_FILE) $(CLIENT_PATH)
	cd $(CLIENT_PATH) && \
	flogo create --cv $(FLOGO_VER) -f $(CLIENT_FILE) $(CLIENT_NAME)
	cp $(PATCH_PATH)/codegen.sh $(CLIENT_PATH)/$(CLIENT_NAME)
	cd $(CLIENT_PATH)/$(CLIENT_NAME) && ./codegen.sh $(FE_HOME) && chmod +x src/gomodedit.sh

.PHONY: build-client
build-client: $(CLIENT_PATH)/$(CLIENT_NAME)
	cd $(CLIENT_PATH)/$(CLIENT_NAME)/src && ./gomodedit.sh
	-cd $(CLIENT_PATH)/$(CLIENT_NAME) && flogo build -e

.PHONY: run
run:
	FLOGO_APP_PROP_RESOLVERS=env FLOGO_APP_PROPS_ENV=auto PORT=8989 FLOGO_LOG_LEVEL=DEBUG FLOGO_SCHEMA_SUPPORT=true FLOGO_SCHEMA_VALIDATION=false CRYPTO_PATH=$(FAB_PATH)/first-network/crypto-config $(CLIENT_PATH)/$(CLIENT_NAME)/bin/$(CLIENT_NAME)

.PHONY: create-fe-client
create-fe-client: $(FE_CLIENT_FILE) clean-client
	mkdir -p $(CLIENT_PATH)
	cp $(FE_CLIENT_FILE) $(CLIENT_PATH)
	cd $(CLIENT_PATH) && \
	flogo create --cv $(FLOGO_VER) -f $(FE_CLIENT_FILE) $(CLIENT_NAME)
	cp $(PATCH_PATH)/codegen.sh $(CLIENT_PATH)/$(CLIENT_NAME)
	cd $(CLIENT_PATH)/$(CLIENT_NAME) && ./codegen.sh $(FE_HOME) && chmod +x src/gomodedit.sh
