#Plugin information
PLUGIN_NAME := terraform-provider-example
PLUGIN_DIR := pkg/bin
PLUGIN_PATH := $(PLUGIN_DIR)/$(PLUGIN_NAME)
DIST_DIR := pkg/dist
VERSION := 0.0.3

TEST?=$$(go list ./... |grep -v 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

all: build test testacc

#Build
build:
	env GOOS=linux   GOARCH=amd64 go build -ldflags '-s -w' -o $(PLUGIN_DIR)/linux/$(PLUGIN_NAME) || exit 1
	env GOOS=linux   GOARCH=386   go build -ldflags '-s -w' -o $(PLUGIN_DIR)/linux86/$(PLUGIN_NAME) || exit 1
	env GOOS=darwin  GOARCH=amd64 go build -ldflags '-s -w' -o $(PLUGIN_DIR)/darwin/$(PLUGIN_NAME) || exit 1
	env GOOS=darwin  GOARCH=386   go build -ldflags '-s -w' -o $(PLUGIN_DIR)/darwin86/$(PLUGIN_NAME) || exit 1
	env GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o $(PLUGIN_DIR)/windows/$(PLUGIN_NAME).exe || exit 1
	env GOOS=windows GOARCH=386   go build -ldflags '-s -w' -o $(PLUGIN_DIR)/windows86/$(PLUGIN_NAME).exe || exit 1
	chmod +x $(PLUGIN_DIR)/*

compress:
	mkdir -p $(DIST_DIR)
	rm -f $(DIST_DIR)/*
	zip -j "${CURRENT_DIR}/$(DIST_DIR)/${PLUGIN_NAME}_${VERSION}_linux.zip" "$(PLUGIN_DIR)/linux/$(PLUGIN_NAME)" || exit 1
	zip -j "${CURRENT_DIR}/$(DIST_DIR)/${PLUGIN_NAME}_${VERSION}_linux86.zip" "$(PLUGIN_DIR)/linux86/$(PLUGIN_NAME)" || exit 1
	zip -j "${CURRENT_DIR}/$(DIST_DIR)/${PLUGIN_NAME}_${VERSION}_darwin.zip" "$(PLUGIN_DIR)/darwin/$(PLUGIN_NAME)" || exit 1
	zip -j "${CURRENT_DIR}/$(DIST_DIR)/${PLUGIN_NAME}_${VERSION}_darwin86.zip" "$(PLUGIN_DIR)/darwin86/$(PLUGIN_NAME)" || exit 1
	zip -j "${CURRENT_DIR}/$(DIST_DIR)/${PLUGIN_NAME}_${VERSION}_windows.zip" "$(PLUGIN_DIR)/windows/$(PLUGIN_NAME).exe" || exit 1
	zip -j "${CURRENT_DIR}/$(DIST_DIR)/${PLUGIN_NAME}_${VERSION}_windows86.zip" "$(PLUGIN_DIR)/windows86/$(PLUGIN_NAME).exe" || exit 1


dev:
	rm -fv terraform.tfstate*
	rm -fv $(PLUGIN_NAME)
	go build -o $(PLUGIN_NAME)
	terraform init
	terraform apply -auto-approve
	terraform state show user_user1.password
	terraform output user_user1_password > /tmp/password
	cat /tmp/password

test:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS)