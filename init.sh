#!/bin/bash
source /.env

mkdir -p /var/log/http/ \
	/var/www/$appName/ 	\
	&& touch /var/log/http/access.log \
	/var/log/http/error.log

go get "$appPath"
go install "$appPath"
site_pid=$(pidof site)
if [[ ! -z "$site_pid" ]] 
	then
	kill -9 "$site_pid"
fi
site > /var/log/http/access.log 2> /var/log/http/error.log 