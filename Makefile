.PHONY: vendor
vendor:
		@go mod tidy
		@go mod vendor
		@modvendor -copy="**/*.c **/*.h" -v
