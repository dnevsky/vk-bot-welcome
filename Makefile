.PHONY: build run push

build:
	docker build --tag dnevsky/vk-bot-welcome .

run:
	docker run -d -i --rm --env VK_TOKEN=<token> --name vk-bot-welcome -it dnevsky/vk-bot-welcome

push:
	docker push dnevsky/vk-bot-welcome:latest