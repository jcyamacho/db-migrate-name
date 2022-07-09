TOOLS := $(CURDIR)/.tools

tools:
	mkdir -p ${TOOLS}
	GOPATH=${TOOLS} go install github.com/caarlos0/svu@latest
	GOPATH=${TOOLS} go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint:
	${TOOLS}/bin/golangci-lint run

release: tools
	git tag "$(${TOOLS}/bin/svu next)"
	git push --tags
