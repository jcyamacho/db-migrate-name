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
VERSION := $(shell ${TOOLS}/bin/svu $(ARG))
release: tools
	git tag "$(VERSION)"
	git push origin "$(VERSION)"

release-major: ARG=major
release-major: release

release-minor: ARG=minor
release-minor: release

release-patch: ARG=patch
release-patch: release
