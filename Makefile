all: build-container run

build-container:
	docker build -t google-stack .

run: build-container
	docker run -itd --rm google-stack
