GO_CMD=go
GO_LINT_CMD=golint
GITHUB_RELEASE_COMMAND=github-release
DOCKER_CMD=docker

VERSION=$(shell cat VERSION)
PRERELEASE=$(cat PRERELEASE)
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_ORIGIN=$(shell git remote get-url origin)
GIT_REPO=$(lastword $(subst :, ,$(GIT_ORIGIN)))
GITHUB_USER=$(firstword $(subst /, ,$(GIT_REPO)))
GITHUB_REPO=$(firstword $(subst ., ,$(lastword $(subst /, ,$(GIT_REPO)))))
GIT_TAG=v$(VERSION)
RELEASE_NAME=$(GIT_TAG)
RELEASE_NOTES_FILE=RELEASE_NOTES.md

GITHUB_RELEASE_FLAGS=--user stratumn --repo go --tag '$(GIT_TAG)'
GITHUB_RELEASE_RELEASE_FLAGS=$(GITHUB_RELEASE_FLAGS) --name '$(RELEASE_NAME)' --description "$$(cat $(RELEASE_NOTES_FILE))"

GO_LIST=$(GO_CMD) list
GO_BUILD=$(GO_CMD) build -ldflags '-X main.version=$(VERSION) -X main.commit=$(GIT_COMMIT)'
GO_TEST=$(GO_CMD) test
GO_LINT=$(GO_LINT_CMD) -set_exit_status
GITHUB_RELEASE_RELEASE=$(GITHUB_RELEASE_COMMAND) release $(GITHUB_RELEASE_RELEASE_FLAGS)
GITHUB_RELEASE_UPLOAD=$(GITHUB_RELEASE_COMMAND) upload $(GITHUB_RELEASE_FLAGS)
GITHUB_RELEASE_EDIT=$(GITHUB_RELEASE_COMMAND) edit $(GITHUB_RELEASE_RELEASE_FLAGS)
DOCKER_BUILD=$(DOCKER_CMD) build
DOCKER_PUSH=$(DOCKER_CMD) push

DIST_DIR=dist
PACKAGES=$(shell $(GO_LIST) ./cmd/...)
OS_ARCHS=darwin_amd64 linux_amd64 windows_amd64
DOCKER_USER=$(GITHUB_USER)
DOCKER_FILE_TEMPLATE=Dockerfile.tpl

TMP_DIR := $(shell mktemp -d)

COVERAGE_LIST=$(shell $(GO_LIST) ./... | grep -v vendor | grep -v testutil | grep -v testcases)
BUILD_LIST=$(foreach package, $(PACKAGES), $(foreach os_arch, $(OS_ARCHS), build_$(package)_$(os_arch)))
ZIP_LIST=$(foreach package, $(PACKAGES), $(foreach os_arch, $(OS_ARCHS), zip_$(package)_$(os_arch)))
DOCKER_FILE_LIST=$(foreach package, $(PACKAGES), docker_file_$(package))
DOCKER_IMAGE_LIST=$(foreach package, $(PACKAGES), docker_image_$(package))
DOCKER_PUSH_LIST=$(foreach package, $(PACKAGES), docker_push_$(package))
GITHUB_UPLOAD_LIST=$(foreach package, $(PACKAGES), $(foreach os_arch, $(OS_ARCHS), github_upload_$(package)_$(os_arch)))

PACKAGE=$(firstword $(subst _, ,$*))
COMMAND=$(lastword $(subst /, ,$(PACKAGE)))
OS=$(word 2, $(subst _, ,$*))
ARCH=$(word 3, $(subst _, ,$*))
OUT_OS_ARCH_DIR=$(DIST_DIR)/$(OS)-$(ARCH)
OUT=$(OUT_OS_ARCH_DIR)/$(COMMAND)
TMP_OS_ARCH_DIR=$(TMP_DIR)/$(OS)-$(ARCH)
TMP_ZIP_DIR=$(TMP_OS_ARCH_DIR)/$(COMMAND)
DOCKER_IMAGE=$(DOCKER_USER)/$(COMMAND)
DOCKER_FILE=$(DIST_DIR)/Dockerfile.$(COMMAND)

.PHONY: $(BUILD_LIST)

all: build

test:
	@echo "==> Running tests"
	$(GO_TEST) ./...

lint:
	@echo "==> Running linter"
	$(GO_LINT) ./...

coverage:
	@echo "" > coverage.txt
	@for d in $(COVERAGE_LIST); do \
	    go test -coverprofile=profile.out -covermode=atomic $$d; \
	    if [ -f profile.out ]; then \
	        cat profile.out >> coverage.txt; \
	        rm profile.out; \
	    fi \
	done

clean:
	@echo "==> Cleaning up"
	rm -rf $(DIST_DIR)

build: $(BUILD_LIST)
zip: $(ZIP_LIST)

git_tag:
	@echo "==> Creating git tag"
	git tag $(GIT_TAG) 2>/dev/null
	git push origin --tags

github_draft:
	@echo "==> Creating Github draft release"
	@if [[ $prerelease != "false" ]]; then \
		$(GITHUB_RELEASE_RELEASE) --draft --pre-release; \
	else \
		$(GITHUB_RELEASE_RELEASE) --draft; \
	fi

github_upload: $(GITHUB_UPLOAD_LIST)

github_publish:
	@echo "==> Publishing Github release"
	@if [[ $prerelease != "false" ]]; then \
		$(GITHUB_RELEASE_EDIT) --pre-release; \
	else \
		$(GITHUB_RELEASE_EDIT); \
	fi

docker_files: $(DOCKER_FILE_LIST)

docker_images: $(DOCKER_IMAGE_LIST)

docker_push: $(DOCKER_PUSH_LIST)

release: test lint clean build zip git_tag github_draft github_upload github_publish docker_files docker_images docker_push

$(BUILD_LIST): build_%:
	@echo "==> Building" $(COMMAND) $(OS) $(ARCH)
	GOOS=$(OS) GOARCH=$(ARCH) $(GO_BUILD) -o $(OUT) $(PACKAGE)

$(ZIP_LIST): zip_%:
	@echo "==> Zipping" $(COMMAND) $(OS) $(ARCH)
	mkdir -p $(TMP_ZIP_DIR)
	cp $(OUT) $(TMP_ZIP_DIR)
	cp LICENSE $(TMP_ZIP_DIR)
	cp RELEASE_NOTES.md $(TMP_ZIP_DIR)
	cp CHANGE_LOG.md $(TMP_ZIP_DIR)
	cd $(TMP_OS_ARCH_DIR) && zip -r $(COMMAND){.zip,} 1>/dev/null
	cp $(TMP_ZIP_DIR).zip $(OUT_OS_ARCH_DIR)

$(GITHUB_UPLOAD_LIST): github_upload_%:
	@echo "==> Uploading Github release file" $(COMMAND)-$(OS)-$(ARCH).zip
	$(GITHUB_RELEASE_UPLOAD) --file $(OUT).zip --name $(COMMAND)-$(OS)-$(ARCH).zip

$(DOCKER_FILE_LIST): docker_file_%:
	@echo "==> Create Dockerfile for" $(DOCKER_IMAGE)
	sed 's/{{CMD}}/$(COMMAND)/g' $(DOCKER_FILE_TEMPLATE) > $(DOCKER_FILE)

$(DOCKER_IMAGE_LIST): docker_image_%:
	@echo "==> Create Docker image" $(DOCKER_IMAGE)
	$(DOCKER_BUILD) -f $(DOCKER_FILE) -t $(DOCKER_IMAGE):$(VERSION) -t $(DOCKER_IMAGE):latest .

$(DOCKER_PUSH_LIST): docker_push_%:
	@echo "==> Pushing Docker image" $(DOCKER_IMAGE)
	$(DOCKER_PUSH) $(DOCKER_IMAGE):$(VERSION)
	$(DOCKER_PUSH) $(DOCKER_IMAGE):latest