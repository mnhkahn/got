.PHONY: install
install:
	go install $(shell go list ./...)