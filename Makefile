buildbuilder:
	docker build -t "nekkkkitch/docker" -f .\Dockerfile .
stop:
	docker-compose stop \
	&& docker-compose rm \
	&& sudo rm -rf pgdata
start:
	docker-compose build --no-cache \
	&& docker-compose up -d
buildauthpb:
	