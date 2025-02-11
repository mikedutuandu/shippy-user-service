build:
	protoc -I. --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/auth/auth.proto
	docker build -t shippy-user-service .

run:
	docker run --net="host" \
		-p 50051 \
		-e DB_HOST=localhost \
		-e DB_PASS=password \
		-e DB_USER=postgres \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		shippy-user-service

