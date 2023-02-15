# References:
# https://seisman.github.io/how-to-write-makefile/index.html

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


# Print help information by default.
.DEFAULT_GOAL := help

##  cli: Install/Update to the latest CLI tool.
.PHONY: cli
cli:
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install -y && \
	rm ./gf


# Check and install CLI tool.
.PHONY: cli.install
cli.install:
	@set -e; \
	gf -v > /dev/null 2>&1 || if [[ "$?" -ne "0" ]]; then \
  		echo "GoFame CLI is not installed, start proceeding auto installation..."; \
		make cli; \
	fi;

##  dao: Generate Go files for Dao/Do/Entity.
.PHONY: dao
dao: cli.install
	@echo "******** gf gen dao ********"
	@GF_GCFG_FILE=config.yaml gf gen dao

##  service: Generate Go files for Service.
.PHONY: service
service: cli.install
	@echo "******** gf gen service ********"
	@gf gen service

##  run: Run gf2-demo-api for development environment.
.PHONY: run
run: cli.install dao service
	@echo "******** gf run ${APISERVER_PATH} ********"
	@gf run ${APISERVER_PATH}

##  run.cli: Run gf2-demo-cli for development environment.
.PHONY: run.cli
run.cli: cli.install dao service
	@echo "******** gf run ${CLI_PATH} ********"
	@gf run ${CLI_PATH}

##  build: Build gf2-demo-api binary.
.PHONY: build
build: cli.install dao service
	@echo "******** gf build ${APISERVER_PATH} ********"
	@${SED} -i '/^      version:/s/version:.*/version: ${VERSION}/' hack/config.yaml
	@gf build ${APISERVER_PATH}

##  build.cli: Build gf2-demo-cli binary.
.PHONY: build.cli
build.cli: cli.install dao service
	@echo "******** gf build ${CLI_PATH} ********"
	@${SED} -i '/^      version:/s/version:.*/version: ${VERSION}/' hack/config.yaml
	@gf build ${CLI_PATH}

# Build image, deploy image and yaml to current kubectl environment and make port forward to local machine.
.PHONY: start
start:
	@set -e; \
	make image; \
	make deploy; \
	make port;

# Build docker image.
.PHONY: image
image: cli.install
	$(eval _TAG  = $(shell git log -1 --format="%cd.%h" --date=format:"%Y%m%d%H%M%S"))
ifneq (, $(shell git status --porcelain 2>/dev/null))
	$(eval _TAG  = $(_TAG).dirty)
endif
	$(eval _TAG  = $(if ${TAG},  ${TAG}, $(_TAG)))
	$(eval _PUSH = $(if ${PUSH}, ${PUSH}, ))
	@gf docker -p -b "-a amd64 -s linux -p temp" -tn $(DOCKER_NAME):${_TAG};


# Build docker image and automatically push to docker repo.
.PHONY: image.push
image.push:
	@make image PUSH=-p;


# Deploy image and yaml to current kubectl environment.
.PHONY: deploy
deploy:
	$(eval _TAG = $(if ${TAG},  ${TAG}, develop))

	@set -e; \
	mkdir -p $(ROOT_DIR)/temp/kustomize;\
	cd $(ROOT_DIR)/manifest/deploy/kustomize/overlays/${_TAG};\
	kustomize build > $(ROOT_DIR)/temp/kustomize.yaml;\
	kubectl   apply -f $(ROOT_DIR)/temp/kustomize.yaml; \
	kubectl   patch -n $(NAMESPACE) deployment/$(DEPLOY_NAME) -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"$(shell date +%s)\"}}}}}";


##  help: Show this help.
.PHONY: help
help: Makefile
	@echo "\nUsage: \n\n    make [TARGETS] \n\nTargets:\n"
	@${SED} -n 's/^##//p' $< | column -t -s ':' | ${SED} -e 's/^/ /'
