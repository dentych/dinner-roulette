# Build app
FROM node:alpine

WORKDIR /work
COPY . /work

RUN npm install
RUN npm run build

# Build real image
FROM nginx

COPY --from=0 /work/dist/ /www/
COPY docker/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
