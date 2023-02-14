CONTAINER_NAME=empty-container
IMAGE_NAME=busybox:latest

clean:
	@docker rm -f $(CONTAINER_NAME) || true && \
	docker rmi -f $(IMAGE_NAME) || true
