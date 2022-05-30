build:
	docker-compose up -d

update:
	docker-compose down
	docker-compose build
	docker-compose up -d

delete:
	docker rm -f $(docker ps -a -q)
	docker volume rm $(docker volume ls -q)