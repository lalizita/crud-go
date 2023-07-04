build:
	docker build -t crud-tasks .

run: build
	docker run --publish 8081:8081 crud-tasks