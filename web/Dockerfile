FROM node:16.13.2

WORKDIR /web/
COPY . .

RUN npm install  --registry=https://registry.npm.taobao.org
RUN npm run build

FROM nginx:alpine

COPY .docker-compose/nginx/conf.d /etc/nginx/conf.d

COPY --from=0 /web/dist /usr/share/nginx/html
# COPY ./dist /usr/share/nginx/html
RUN cat /etc/nginx/nginx.conf
RUN cat /etc/nginx/conf.d/my.conf
RUN ls -al /usr/share/nginx/html
