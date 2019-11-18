# -------------
# VARIABLES
# -------------

# -------------
# FUNCTIONS
# -------------


# -------------
# TASKS
# -------------
.PHONY: fmt
fmt:
	@gofmt -w -s -d configuration
	@gofmt -w -s -d main.go

.PHONY: test-build
test-build: fmt
	@go build

.PHONY: build
build:
	@docker build --tag $$(basename $$(git rev-parse --show-toplevel)):$$(date +%s) --build-arg "BRANCH=$(BRANCH)" build/docker
