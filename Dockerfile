FROM golang:1.8

LABEL maintainer="Stephen Afam-Osemene <stephenafamo@gmail.com>"

EXPOSE 80 443

RUN mkdir -p /var/log/http/ \
	/var/www/stephenafamo/ \
	&& touch /var/log/http/access.log /var/log/http/error.log
