run:
	docker build -t api-image .
	docker run -d -p 8000:8000 --name api api-image
	docker logs --follow api
stop:
	docker stop api 
	docker rm api
	docker rmi api-image
	# docker rmi golang:1.20-alpine
	# docker rmi alpine:latest
	docker image prune -f