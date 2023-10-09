# References:
# https://seisman.github.io/how-to-write-makefile/index.html

# Be same as gf version in go.mod.
GF_VERSION = v2.5.4
GF_PATH = ${HOME}/.gf/${GF_VERSION}
GF_BIN = ${GF_PATH}/gf

export PATH := ${GF_PATH}:${GOROOT/bin}:${PATH}

ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"

APISERVER_NAME = gf2-demo-api
CLI_NAME = gf2-demo-cli

DEPLOY_NAME = ${APISERVER_NAME}
DOCKER_NAME = ${APISERVER_NAME}

APISERVER_PATH = cmd/gf2-demo-api/${APISERVER_NAME}.go
CLI_PATH = cmd/gf2-demo-cli/${CLI_NAME}.go

VERSION = $(shell git describe --tags --always --match='v*')

SED = sed
ifeq ($(shell uname), Darwin)
	SED = gsed
endif


define USAGE_OPTIONS

Options:

    ARCH         The multiple ARCH to build. Default is "amd64";
                 This option is available for: make build*;
                 Example: make build ARCH="amd64,arm64"

    OS           The multiple OS to build. Default is "darwin,linux";
                 This option is available for: make build*;
                 Example: make build OS="linux,darwin,windows"

    TAG          Docker image tag. Default is generated from current git commit;
                 This option is available for: make image*;
                 Example: make image TAG="v0.6.0"
endef

export USAGE_OPTIONS

GF_BUILD_ARGS =
ifneq (${OS},)
	GF_BUILD_ARGS = -s ${OS}
endif

ifneq (${ARCH},)
	GF_BUILD_ARGS += -a ${ARCH}
endif


# Print help information by default
.DEFAULT_GOAL := help

##   cli: Install/Update to the latest Gf Cli tool
.PHONY: cli
cli:
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/download/${GF_VERSION}/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	mkdir -p ${GF_PATH} && \
	mv ./gf ${GF_PATH}


# Check and install CLI tool
.PHONY: cli.install
cli.install:
	@echo "******** install gf cli ********"
	@set -e; \
	if ! ${GF_BIN} -v >/dev/null 2>&1; then \
  		echo "GoFame CLI is not installed, start proceeding auto installation..."; \
		make cli; \
	elif [[ $$(${GF_BIN} -v|grep -i cli|grep -Eio "v[0-9]+\.[0-9]+\.[0-9]+"|head -1) != ${GF_VERSION} ]];then \
  		echo "GoFame CLI version is not equal to ${GF_VERSION}, start proceeding auto installation..."; \
		make cli; \
	else \
		echo "GoFame CLI is already installed and version is right: $(GF_VERSION)"; \
	fi;

# Check and install golangci-lint tool
.PHONY: lint.install
lint.install:
	@set -e; \
	golangci-lint --version >/dev/null 2>&1 || \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

##   lint: Run golangci-lint
.PHONY: lint
lint: lint.install
	@echo "******** golangci-lint run ********"
	@golangci-lint run --timeout=10m

##   ctrl: Generate Go files for Controller
.PHONY: ctrl
ctrl: cli.install
	@echo "******** gf gen ctrl ********"
	@${GF_BIN} gen ctrl -k api/sdk

##   dao: Generate Go files for Dao/Do/Entity
.PHONY: dao
dao: cli.install
	@echo "******** gf gen dao ********"
	@${GF_BIN} gen dao

##   service: Generate Go files for Service
.PHONY: service
service: cli.install
	@echo "******** gf gen service ********"
	@${GF_BIN} gen service

##   run: Run gf2-demo-api for development environment
.PHONY: run
run: ctrl dao service
	@echo "******** gf run ${APISERVER_PATH} ********"
	@${GF_BIN} run ${APISERVER_PATH}

##   run.cli: Run gf2-demo-cli for development environment
.PHONY: run.cli
run.cli: ctrl dao service
	@echo "******** gf run ${CLI_PATH} ********"
	@${GF_BIN} run ${CLI_PATH}

##   build: Build gf2-demo-api binary
.PHONY: build
build: ctrl service
	@echo "******** gf build ${APISERVER_PATH} ********"
	@${SED} -i '/^      version:/s/version:.*/version: ${VERSION}/' hack/config.yaml
	@${GF_BIN} build ${APISERVER_PATH} ${GF_BUILD_ARGS}

##   build.cli: Build gf2-demo-cli binary
.PHONY: build.cli
build.cli: ctrl service
	@echo "******** gf build ${CLI_PATH} ********"
	@${SED} -i '/^      version:/s/version:.*/version: ${VERSION}/' hack/config.yaml
	@${GF_BIN} build ${CLI_PATH} ${GF_BUILD_ARGS}

# Build image, deploy image and yaml to current kubectl environment and make port forward to local machine
.PHONY: start
start:
	@set -e; \
	make image; \
	make deploy; \
	make port;

##   image: Build docker image
.PHONY: image
image: cli.install
	$(eval _TAG  = $(shell git log -1 --format="%cd.%h" --date=format:"%Y%m%d%H%M%S"))
ifneq (, $(shell git status --porcelain 2>/dev/null))
	$(eval _TAG  = $(_TAG).dirty)
endif
	$(eval _TAG  = $(if ${TAG}, ${TAG}, ${_TAG}))
	@echo ${DOCKER_NAME}:${_TAG}
	@docker image build -t ${DOCKER_NAME}:${_TAG} .

##   image.push: Build docker image and automatically push to docker repo
.PHONY: image.push
image.push: image
	@docker image push ${DOCKER_NAME}:${_TAG}

#   Deploy image and yaml to current kubectl environment
.PHONY: deploy
deploy:
	$(eval _TAG = $(if ${TAG},  ${TAG}, develop))

	@set -e; \
	mkdir -p $(ROOT_DIR)/temp/kustomize;\
	cd $(ROOT_DIR)/manifest/deploy/kustomize/overlays/${_TAG};\
	kustomize build > $(ROOT_DIR)/temp/kustomize.yaml;\
	kubectl   apply -f $(ROOT_DIR)/temp/kustomize.yaml; \
	kubectl   patch -n $(NAMESPACE) deployment/$(DEPLOY_NAME) -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"$(shell date +%s)\"}}}}}";


##   help: Show this help
.PHONY: help
help: Makefile
	@echo "\nUsage: \n\n    make [TARGETS] [OPTIONS] \n\nTargets:\n"
	@${SED} -n 's/^##//p' $< | column -t -s ':' | ${SED} -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
