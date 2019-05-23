

all:
	mkdir -p pb/oneoftest
	protoc --go_out=paths=source_relative:./pb/oneoftest oneoftest.proto