#!/bin/sh

set -e

case "$1" in
    *.yaml|*.yml) set -- registry serve "$@" ;;
    serve|garbage-collect|help|-*) set -- registry "$@" ;;
esac

if [ `echo "$@" | grep -c "registry serve"` == 1 ]; then
    exec "$@"& >> /dev/stdout
    # Wait for registry startup.
    sleep 2
    /opt/docker-registry-ui
else
    exec "$@"
fi
