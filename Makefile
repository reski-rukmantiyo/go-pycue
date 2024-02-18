# include .env

PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
#GOPATH="$(GOBASE)/vendor:$(GOBASE)"
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
GOSOURCEFILES="main.go"
SWAGDOCS="./docs"

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# PID file will keep the process id of the server
PID=/tmp/.$(PROJECTNAME).pid

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

build:
	echo "Update Swagger Docs"
	#goreleaser build --snapshot --rm-dist --skip=validate
	#goreleaser release --snapshot --clean
	make swag
	go build -o $(GOBIN)/$(PROJECTNAME) $(GOSOURCEFILES)

swag:
	echo "Create Swagger files"
	swag init -g $(GOSOURCEFILES) -o $(SWAGDOCS)

compile:
	echo "Goreleaser Compile only" 
	goreleaser release --skip-announce --skip-publish --rm-dist --snapshot

publish:
	echo "Goreleaser Release the Apps"
	goreleaser release  --rm-dist
	
all: build compile