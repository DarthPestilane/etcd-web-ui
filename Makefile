image:
	docker build -t darthminion/etcd-web-ui . --platform=linux/amd64 && docker image prune -f

backend-linux:
	cd back_end && $(MAKE) build-linux

frontend:
	cd front_end && $(MAKE)
