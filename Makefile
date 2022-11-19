BUILD_ENVPARAMS:=CGO_ENABLE=0

.PHONY: build
build:
	$(BUILD_ENVPARAMS) go build -o out/ddcsync cmd/ddcsync/main.go
