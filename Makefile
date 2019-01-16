# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOGENERATE=$(GOCMD) generate
GOLIST=$(GOCMD) list

# Git parameters
GITCMD=git
GITCHECKOUT=$(GITCMD) checkout

# Sources parameters
SOURCE_ENTRYPOINT=./cmd
# Binary parameters

BINARY_NAME=dbc4go
BINARY_DESTINATION=./bin
BINARY_PATH=$(BINARY_DESTINATION)/$(BINARY_NAME)

# Tagets
all:	test build

# Build with contracts
buildwc:
		$(GOGENERATE) ./... && $(GOBUILD) -o $(BINARY_PATH) -v $(SOURCE_ENTRYPOINT) && $(GITCHECKOUT) -- .

build:
		$(GOBUILD) -o $(BINARY_PATH) -v $(SOURCE_ENTRYPOINT)
# Unit tests
utest:
		$(GOTEST) --cover `$(GOLIST) ./... | grep -v "/examples"`
clean: 
		$(GOCLEAN) $(SOURCE_ENTRYPOINT)
		rm -f $(BINARY_PATH)
run:
		$(GOBUILD) -o $(BINARY_PATH) -v $(SOURCE_ENTRYPOINT)
		$(BINARY_PATH)&
