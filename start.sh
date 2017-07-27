#!/bin/bash
app="stephenafamo"

docker build -t $app .

docker stop "$app" && docker rm -f "$app"

docker run -d -t \
	--restart=always \
	-p 47852:80 \
    --net=crs_backend \
    --name=$app \
    -v "$PWD"/site:/go/src/github.com/"$app"/site \
    -v "$PWD"/resources:/var/www/stephenafamo \
    -w /go/src/github.com/"$app" \
    $app
