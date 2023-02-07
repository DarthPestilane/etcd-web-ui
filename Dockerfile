FROM nginx:alpine

ENV UI_ENDPOINT='127.0.0.1:2379'
ENV UI_USERNAME=''
ENV UI_PASSWORD=''
ENV UI_KEY_PREFIX='/'

COPY front_end/nginx.conf /etc/nginx/nginx.conf
COPY front_end/dist/ /var/local/www
COPY back_end/bin/api /bin/api
COPY docker-entrypoint.sh /docker-entrypoint.sh

RUN chmod +x /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]

CMD ["nginx", "-g", "daemon off;"]
