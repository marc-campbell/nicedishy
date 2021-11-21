SHELL := /bin/bash -o pipefail
VERSION ?=`git describe --tags`
DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"`
VERSION_PACKAGE = github.com/marc-campbell/nicedishy/pkg/version
GIT_TREE = $(shell git rev-parse --is-inside-work-tree 2>/dev/null)
ifneq "$(GIT_TREE)" ""
define GIT_UPDATE_INDEX_CMD
git update-index --assume-unchanged
endef
define GIT_SHA
`git rev-parse HEAD`
endef
else
define GIT_UPDATE_INDEX_CMD
echo "Not a git repo, skipping git update-index"
endef
define GIT_SHA
""
endef
endif

define LDFLAGS
-ldflags "\
	-X ${VERSION_PACKAGE}.version=${VERSION} \
	-X ${VERSION_PACKAGE}.gitSHA=${GIT_SHA} \
	-X ${VERSION_PACKAGE}.buildTime=${DATE} \
"
endef

CURRENT_USER := $(shell id -u -n)
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

.PHONY: all
all: bin/nicedishy-api bin/generate-data

.PHONY: bin/nicedishy-api
bin/nicedishy-api:
	go build -o bin/nicedishy-api ./cmd/api

.PHONY: bin/generate-data
bin/generate-data:
	go build -o bin/generate-data ./cmd/generate-data

.PHONY: fmt
fmt:
	go fmt ./pkg/... ./cmd/...

.PHONY: vet
vet:
	go vet ./pkg/... ./cmd/...

.PHONY: test
test: fmt vet
	go test ./pkg/... ./cmd/...

.PHONY: thanos
thanos:
	rm -rf kube-thanos
	git clone https://github.com/thanos-io/kube-thanos.git kube-thanos
	cd kube-thanos && make vendor
	cd kube-thanos && ./build.sh ../thanos.jsonnet
	rm -rf ./kustomize/overlays/dev/thanos && mkdir -p ./kustomize/overlays/dev/thanos
	cp -r ../kube-thanos/manifests/* ./kustomize/overlays/dev/thanos
	rm -rf kube-thanos

.PHONY: swag
swag:
	cd cmd/api && swag init