up:
	docker-compose up -d
stop:
	docker-compose stop
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

