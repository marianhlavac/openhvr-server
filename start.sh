#!/usr/bin/env sh

docker run -d --mount source=openhvr-data,destination=/go/src/github.com/mmajko/openhvr-server/_data -p 47023:47023 marianhlavac/openhvr
