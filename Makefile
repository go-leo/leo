
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
# 默认 help
.DEFAULT_GOAL := help


# 版本发布测试
version-dry-run:
	bash ./scripts/version.sh --dry-run=1


# 版本发布
version:
	bash ./scripts/version.sh
