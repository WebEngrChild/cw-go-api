IMAGE_NAME=my-go-app

.PHONY: build run

build:
	docker build -t $(IMAGE_NAME) .

run:
	docker run -p 8080:8080 $(IMAGE_NAME)

deploy:
	copilot svc deploy --name cw-go-api --env test