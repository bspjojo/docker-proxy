FROM golang:1.15 AS build

WORKDIR /go/src/app
COPY . .

RUN mv config/config.prod.json config/config.json

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o proxy-gen

FROM nginx:1.17.8

COPY docker-start-script.sh /app/docker-start-script.sh
COPY config/config.prod.json /app/config/config.json
COPY --from=build /go/src/app/proxy-gen /app/proxy-gen

RUN ls -l /app

RUN mkdir -p /etc/nginx/conf.d/

EXPOSE 80

ENTRYPOINT ["/app/docker-start-script.sh"]