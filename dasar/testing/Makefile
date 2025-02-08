unit-test:
	go test ./...  --short

integration-test:
	# spin up docker redis container
	docker run -d --name redis-stack 6379:6379 -p 8001:8001 redis/redis-stack:latest
	go test ./... 
	# stop redis container
	docker container stop redis-stack
	docker system prune -f