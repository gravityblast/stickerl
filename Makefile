GO_CMD=go
GO_TEST=TRAFFIC_ENV=test $(GO_CMD) test -v
GO_BUILD=$(GO_CMD) build -v

all: build
test: RunTests
build: BuildApp

BuildApp:RunTests
	$(GO_BUILD)

RunTests:
	$(GO_TEST)
