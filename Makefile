PROJECT_NAME := Pulumi k0s Resource Provider

PACK             := k0s
PACKDIR          := sdk
PROJECT          := github.com/ydkn/pulumi-k0s
NODE_MODULE_NAME := @ydkn/pulumi-k0s
NUGET_PKG_NAME   := Pulumi.K0s

PROVIDER      := pulumi-resource-${PACK}
CODEGEN       := pulumi-gen-${PACK}
VERSION       ?= $(shell pulumictl get version)
PROVIDER_PATH := provider
VERSION_PATH  := ${PROVIDER_PATH}/pkg/version.Version

SCHEMA_FILE := provider/cmd/pulumi-resource-k0s/schema.json
GOPATH      := $(shell go env GOPATH)

WORKING_DIR     := $(shell pwd)
TESTPARALLELISM := 4
GO_TEST         := go test -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM}

ensure::
	cd provider && go mod tidy
	cd sdk && go mod tidy
	cd tests && go mod tidy

gen::
	(cd provider && go build -o $(WORKING_DIR)/bin/${CODEGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/$(CODEGEN))

provider::
	(cd provider && VERSION=${VERSION} go generate cmd/${PROVIDER}/main.go)
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

provider_debug::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -gcflags="all=-N -l" -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

test_provider::
	cd provider/pkg && go test -short -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM} ./...

dotnet_sdk:: DOTNET_VERSION := $(shell pulumictl get version --language dotnet)
dotnet_sdk::
	rm -rf sdk/dotnet
	$(WORKING_DIR)/bin/$(CODEGEN) -version=${DOTNET_VERSION} dotnet $(SCHEMA_FILE) $(CURDIR)
	cd ${PACKDIR}/dotnet/&& \
		echo "${DOTNET_VERSION}" >version.txt && \
		dotnet build /p:Version=${DOTNET_VERSION}

go_sdk::
	rm -rf sdk/go
	$(WORKING_DIR)/bin/$(CODEGEN) -version=${VERSION} go $(SCHEMA_FILE) $(CURDIR)

nodejs_sdk:: VERSION := $(shell pulumictl get version --language javascript)
nodejs_sdk::
	rm -rf sdk/nodejs
	$(WORKING_DIR)/bin/$(CODEGEN) -version=${VERSION} nodejs $(SCHEMA_FILE) $(CURDIR)
	jq '.name = "$(NODE_MODULE_NAME)"' ${PACKDIR}/nodejs/package.json > ${PACKDIR}/nodejs/package.json.new
	mv ${PACKDIR}/nodejs/package.json.new ${PACKDIR}/nodejs/package.json
	cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp -r ../../README.md ../../LICENSE ./scripts ./package.json ./yarn.lock ./bin/ && \
		sed 's/$${VERSION}/$(VERSION)/g' ./bin/package.json > /tmp/nodejs-package.json && \
		mv /tmp/nodejs-package.json ./bin/package.json

python_sdk:: PYPI_VERSION := $(shell pulumictl get version --language python)
python_sdk::
	rm -rf sdk/python
	$(WORKING_DIR)/bin/$(CODEGEN) -version=${VERSION} python $(SCHEMA_FILE) $(CURDIR)
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py > /tmp/python-setup.py && \
		mv /tmp/python-setup.py ./bin/setup.py && \
		cd ./bin && python3 setup.py build sdist

.PHONY: build
build:: gen provider dotnet_sdk go_sdk nodejs_sdk python_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

lint::
	for DIR in "provider" "sdk" "tests" ; do \
		pushd $$DIR && golangci-lint run -c ../.golangci.yml --timeout 10m && popd ; \
	done

install:: install_nodejs_sdk install_dotnet_sdk
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

test_all::
	cd provider/pkg && $(GO_TEST) ./...
	cd tests/sdk/nodejs && $(GO_TEST) ./...
	cd tests/sdk/python && $(GO_TEST) ./...
	cd tests/sdk/dotnet && $(GO_TEST) ./...
	cd tests/sdk/go && $(GO_TEST) ./...

install_dotnet_sdk::
	rm -rf $(WORKING_DIR)/nuget/$(NUGET_PKG_NAME).*.nupkg
	mkdir -p $(WORKING_DIR)/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

install_python_sdk::
	#target intentionally blank

install_go_sdk::
	#target intentionally blank

install_nodejs_sdk::
	-yarn unlink --cwd $(WORKING_DIR)/sdk/nodejs/bin
	yarn link --cwd $(WORKING_DIR)/sdk/nodejs/bin

.PHONY: publish_sdks
publish_sdks:: publish_python_sdk publish_go_sdk publish_nodejs_sdk

publish_python_sdk:: python_sdk
	cd $(WORKING_DIR)/sdk/python/bin && \
		twine upload -u __token__ -p ${PYPI_TOKEN} dist/*

publish_go_sdk:: go_sdk
	#target intentionally blank

publish_nodejs_sdk:: nodejs_sdk
	cd $(WORKING_DIR)/sdk/nodejs/bin && \
		yarn publish --access=public --no-git-tag-version --new-version="$(VERSION)"

.PHONY: publish_plugin
publish_plugin:
	GOOS=linux GOARCH=amd64 $(MAKE) publish_plugin_arch
	GOOS=linux GOARCH=arm64 $(MAKE) publish_plugin_arch
	GOOS=darwin GOARCH=amd64 $(MAKE) publish_plugin_arch
	GOOS=darwin GOARCH=arm64 $(MAKE) publish_plugin_arch
	GOOS=windows GOARCH=amd64 $(MAKE) publish_plugin_arch

publish_plugin_arch:
	$(MAKE) provider
	mkdir -p dist
	tar -cf dist/pulumi-resource-k0s-v$(VERSION)-$(GOOS)-$(GOARCH).tar README.md LICENSE
	cd bin && tar -rf ../dist/pulumi-resource-k0s-v$(VERSION)-$(GOOS)-$(GOARCH).tar pulumi-resource-k0s
	gzip -9 dist/pulumi-resource-k0s-v$(VERSION)-$(GOOS)-$(GOARCH).tar
