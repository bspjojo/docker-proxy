FROM golang:1.15

WORKDIR /go/src/app
COPY . .

RUN mv config/config.prod.json config/config.json

RUN go get -d -v ./...
RUN go install -v ./...

RUN mkdir -p /etc/nginx/conf.d/

EXPOSE 80

CMD ["main"]