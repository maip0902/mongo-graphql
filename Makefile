.PHONY: up
up:
	docker-compose up -d
.PHONY: stop
stop:
	docker-compose stop
.PHONY: down
down:
	docker-compose down
.PHONY: ps
ps:
	docker ps
.PHONY: build
build:
	docker-compose build
.PHONY: app
app:
	docker exec -it go-test bash

