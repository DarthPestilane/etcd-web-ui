#!/bin/sh
# vim:sw=4:ts=4:et

set -ex

nohup /bin/api start --endpoint=$UI_ENDPOINT --username=$UI_USERNAME --password=$UI_PASSWORD --key-prefix=$UI_KEY_PREFIX &

exec "$@"
