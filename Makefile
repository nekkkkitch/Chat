buildbuilder: # call in /Chat
	docker build -t "nekkkkitch/docker" -f .\Dockerfile .
stop:
	docker-compose stop \
	&& docker-compose rm \
	&& sudo rm -rf pgdata
start:
	docker-compose build --no-cache \
	&& docker-compose up -d
buildauthpb: # call in dir /pkg/grpc
	protoc --proto_path=proto/authService --go_out=pb/authService --go-grpc_out=pb/authService proto/authService/*.proto