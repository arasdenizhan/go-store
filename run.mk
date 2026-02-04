docker_build:
	docker build -t go-store:latest .

docker_run:
	docker run --name=go-store -p 8080:8080 go-store:latest