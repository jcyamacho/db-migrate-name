TOOLS := $(CURDIR)/.tools

tools:
	mkdir -p ${TOOLS}
	GOPATH=${TOOLS} go install github.com/caarlos0/svu@latest
	GOPATH=${TOOLS} go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	GOPATH=${TOOLS} go install github.com/goreleaser/goreleaser@latest

lint: tools
	${TOOLS}/bin/golangci-lint run

build: tools
	goreleaser release --snapshot --rm-dist

ARG ?= next
release: tools
	git tag "$(${TOOLS}/bin/svu $(ARG))"
	git push --tags

release-major: ARG=major
release-major: release

release-minor: ARG=minor
release-minor: release

release-patch: ARG=patch
release-patch: release
