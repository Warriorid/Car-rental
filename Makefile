
rebuild: down build up

build:
	docker-compose build --no-cache

refresh: down
	docker-compose up --build -d

up:
	docker-compose up

down:
	docker-compose down
