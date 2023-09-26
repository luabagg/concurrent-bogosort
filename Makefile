# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test

# Test
TEST_FLAGS=-race

default: all

all: test

test:
	@echo "Running tests..."
	$(GOTEST) -v $(TEST_FLAGS) ./...

.PHONY: all test