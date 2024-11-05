start:
	docker build -t "build_image" -f .\Dockerfile .
stop:
	docker-compose stop \
	&& docker-compose rm \
	&& sudo rm -rf pgdata