PROJECT_NAME=ratelimitd
GIT_TAG = $(shell git tag | grep ^v | sort -V | tail -n 1)
GIT_REVISION = $(shell git rev-parse --short HEAD)
GIT_SUMMARY = $(shell git describe --tags --dirty --always)
GO_IMPORT_PATH=github.com/bububa/ratelimitd/app
DIST_PATH=./dist
EXEC_PATH=/usr/local/bin
LDFLAGS = -X $(GO_IMPORT_PATH).GitTag=$(GIT_TAG) -X $(GO_IMPORT_PATH).GitRevision=$(GIT_REVISION) -X $(GO_IMPORT_PATH).GitSummary=$(GIT_SUMMARY) -s -w -extldflags "-static" # -X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn

.PHONY : all

all: deploy 

proto: 
	protoc --gofast_out=$(GOPATH)/src ./pb/ratelimitd/*.proto

server:
ifeq (,$(wildcard $(DIST_PATH)/$(PROJECT_PATH)))
	rm -rf $(DIST_PATH)/$(PROJECT_NAME)
endif
	go build -ldflags "$(LDFLAGS)" -o $(DIST_PATH)/$(PROJECT_NAME) 

clean:
	rm -rf $(DIST_PATH)/$(PROJECT_NAME)

install:
	sudo mv $(DIST_PATH)/$(PROJECT_NAME) $(EXEC_PATH)/$(PROJECT_NAME);
	sudo chown root:root $(EXEC_PATH);
	sudo chmod 755 $(EXEC_PATH);

deploy: 
	make server
	sudo make install

