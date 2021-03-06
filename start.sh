#!/bin/bash
source ./.env

docker build -t $appName .

docker stop "$appName" && docker rm -f "$appName"

docker run -d -t \
	--restart=always \
	-p "$appPort":80 \
    --net="$dockerNetwork" \
    --name=$appName \
    -v "$PWD"/site:/go/src/"$appPath" \
    -v "$PWD"/resources:"$resourcesPath" \
    -w /go/src/"$appPath" \
    $appName
