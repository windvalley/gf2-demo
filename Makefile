ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "gf2-demo-api"
DOCKER_NAME = "gf2-demo-api"

SED = gsed

APISERVER_CMD = cmd/gf2-demo-api/gf2-demo-api.go
JOB_CMD = cmd/gf2-demo-job/gf2-demo-job.go

VERSION = $(shell git describe --tags --always --match='v*')


# Install/Update to the latest CLI tool.
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

# Generate Go files for DAO/DO/Entity.
.PHONY: dao
dao: cli.install
	@echo "******** gf gen dao ********"
	@gf gen dao

# Generate Go files for Service.
.PHONY: service
service: cli.install
	@echo "******** gf gen service ********"
	@gf gen service

# Run apiserver for development environment.
.PHONY: run
run: cli.install dao service
	@echo "******** gf run ${APISERVER_CMD} ********"
	@gf run ${APISERVER_CMD}

# Run cronjob for development environment.
.PHONY: run.job
run.job: cli.install dao service
	@echo "******** gf run ${JOB_CMD} ********"
	@gf run ${JOB_CMD}

# Build apiserver binary.
.PHONY: build
build: cli.install dao service
	@echo "******** gf build ${APISERVER_CMD} ********"
	@${SED} -i '/^      version:/s/version:.*/version: ${VERSION}/' hack/config.yaml
	@gf build ${APISERVER_CMD}

# Build cronjob binary.
.PHONY: build.job
build.job: cli.install dao service
	@echo "******** gf build ${JOB_CMD} ********"
	@${SED} -i '/^      version:/s/version:.*/version: ${VERSION}/' hack/config.yaml
	@gf build ${JOB_CMD}

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


