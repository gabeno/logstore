compile:
	protoc --go_out=. --go_opt=paths=source_relative --proto_path=. api/v1/log.proto 

test:
	go test -race ./..
