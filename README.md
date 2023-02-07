# ETCD WEB UI

**ONLY FOR ETCD API V3**

## Introduction

This is a ETCD web UI, featured with basic operations of etcd kv.

## Usage

```sh
docker run \
    -p 8283:80 \
    --env=UI_ENDPOINT=127.0.0.1:2379 \
    --env=UI_USERNAME=root \
    --env=UI_PASSWORD=secret \
    --env=UI_KEY_PREFIX=/prefix/ \
    darthminion/etcd-web-ui
```
