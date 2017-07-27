#!/bin/bash
docker exec stephenafamo go get github.com/stephenafamo/site
docker exec stephenafamo go install github.com/stephenafamo/site
site_pid=$(docker exec stephenafamo pidof site)
if [[ ! -z "$site_pid" ]] 
	then
	docker exec stephenafamo kill -9 "$site_pid"
fi
docker exec stephenafamo bash -c "nohup site > /var/log/http/access.log 2> /var/log/http/error.log &"