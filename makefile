.PHONY: build run

build:
	docker build -t salt:dev .

run:
	docker run -t salt:dev
