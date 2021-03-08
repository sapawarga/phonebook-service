.PHONY: clean build packing

clean: 
	rm -rf ./main
	rm -vf ./transport/grpc/phonebook/*.pb.go

proto: clean
	@echo "--- Preparing proto output directories ---"
	@mkdir -p ./transport/grpc/phonebook
	@cd ./transport/grpc/phonebook && protoc *.proto --go_out=plugins=grpc:.
	@echo "--- Finished generate proto file ---"
	
build:
	@GOOS=linux GOARCH=amd64
	@echo ">> Building GRPC..."
	@go build -o phonebook-grpc ./cmd/grpc
	@echo ">> Finished"
	