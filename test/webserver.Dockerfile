FROM alpine:3.10.3
RUN apk add lighttpd lighttpd-mod_auth
COPY lorem_ipsum.html /var/www/localhost/htdocs/index.html
COPY data.json /var/www/localhost/htdocs/
COPY htpasswd /etc/lighttpd/
COPY lighttpd.conf /etc/lighttpd/

CMD ["lighttpd", "-D", "-f", "/etc/lighttpd/lighttpd.conf"]
