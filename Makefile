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

.PHONY: build
build: bin/nicedishy-api bin/generate-data bin/grafana-proxy bin/nicedishy-dev

.PHONY: run
run: bin/nicedishy-api
	./bin/nicedishy-api api

.PHONY: run-grafana-proxy
run-grafana-proxy: bin/grafana-proxy
	./bin/grafana-proxy run

.PHONY: bin/nicedishy-api
bin/nicedishy-api:
	go build -o bin/nicedishy-api ./cmd/api

.PHONY: bin/nicedishy-dev
bin/nicedishy-dev:
	go build -o bin/nicedishy-dev ./cmd/dev

.PHONY: bin/grafana-proxy
bin/grafana-proxy:
	go build -o bin/grafana-proxy ./cmd/dashboard

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

.PHONY: sealedsecrets
sealedsecrets:
	kubeseal --cert ./sealedsecret.pem --format yaml < ./kustomize/overlays/production/google-secret-raw.yaml > ./kustomize/overlays/production/google-secret.yaml
	kubeseal --cert ./sealedsecret.pem --format yaml < ./kustomize/overlays/production/imagepullsecret-raw.yaml > ./kustomize/overlays/production/imagepullsecret.yaml
	kubeseal --cert ./sealedsecret.pem --format yaml < ./migrations/kustomize/overlays/production/timescale-secret-raw.yaml > ./migrations/kustomize/overlays/production/timescale-secret.yaml
