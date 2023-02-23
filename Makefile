GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOINSTALL=$(GOCMD) install
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOGENERATE=$(GOCMD) generate
DOCKER=docker
DOCKER_BUILD=$(DOCKER) build
DOCKER_RUN=$(DOCKER) run

PRJ_NAME=try-golang-extlib
GITHUB_USER=devlights
PKG_NAME=github.com/$(GITHUB_USER)/$(PRJ_NAME)
CMD_PKG=$(PKG_NAME)/cmd/trygolang-extlib

EXAMPLE=""

RM_CMD=rm -f
BIN_NAME=trygolang-extlib
BIN_DIR=bin

.PHONY: all
all: clean build test

.PHONY: prepare
prepare: \
	_go_get \
	_download_sqlite3_database

_go_get:
	$(GOCMD) get -d ./...

_download_sqlite3_database:
	@if [ ! -e "chinook.db" ]; then\
		wget https://www.sqlitetutorial.net/wp-content/uploads/2018/03/chinook.zip;\
		unzip -o chinook.zip;\
		rm -f chinook.zip;\
	fi

.PHONY: build
build: prepare
	$(GOBUILD) -o $(BIN_DIR)/$(BIN_NAME) $(CMD_PKG)

.PHONY: test
test: prepare
	$(GOTEST) -v ./...

.PHONY: clean
clean: prepare
	$(GOCLEAN) -i $(CMD_PKG)
	$(RM_CMD) $(BIN_DIR)/$(BIN_NAME)

.PHONY: install
install: prepare
	$(GOINSTALL) $(BIN_DIR)

.PHONY: run
run: prepare
	$(GORUN) $(CMD_PKG) -onetime -example ${EXAMPLE}

.PHONY: generate
generate: prepare
	$(GOGENERATE) -x ./...
