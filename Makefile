buildbuilder: # call in /Chat
	docker build -t "nekkkkitch/docker" -f .\Dockerfile .
stop:
	docker-compose stop \
	&& docker-compose rm \
	&& sudo rm -rf pgdata
start:
	docker-compose build --no-cache \
	&& docker-compose up -d
buildauthpb: 
	protoc --proto_path=pkg/grpc/proto/authService --go_out=pkg/grpc/pb/authService --go-grpc_out=pkg/grpc/pb/authService pkg/grpc/proto/authService/*.proto
buildmsgpb: 
	protoc --proto_path=pkg/grpc/proto/msgService --go_out=pkg/grpc/pb/msgService --go-grpc_out=pkg/grpc/pb/msgService pkg/grpc/proto/msgService/*.proto