FROM nginx:1.25.4-alpine

COPY .nginx/default.conf /etc/nginx/conf.d/
COPY www/index.html /usr/share/nginx/html/