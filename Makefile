.PHONY: all install

all:
	@rm -rf ./bin
	go build -o ./bin/drawsaface

install:
	go mod tidy