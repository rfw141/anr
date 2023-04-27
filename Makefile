.PHONY: push
# push
push:
	git add .
	git commit -m "ok"
	git push origin master

.PHONY: tidy
# tidy
tidy:
	go mod tidy

.PHONY: buf
# generate proto
buf:
	./bin/buf generate proto

		
.PHONY: generate
# generate client code
generate:
	go generate ./...

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
