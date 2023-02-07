image:
	docker build -t darthminion/etcd-web-ui . && docker image prune -f

backend-linux:
	cd back_end && $(MAKE) build-linux

frontend:
	cd front_end && $(MAKE)
