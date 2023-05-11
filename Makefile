.PHONY: build run push shutdown

build:
	docker build --tag dnevsky/vk-bot-welcome .

run:
	docker-compose up -d vk-bot-welcome

shutdown:
	docker-compose down

push:
	docker push dnevsky/vk-bot-welcome:latest