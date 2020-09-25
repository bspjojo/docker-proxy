# Docker proxy

## What is this?

A simple proxy container that proxies requests through to other containers.

It generates a configuration file for an interanl instance of nginx

## Why is this a thing?

I got annoyed at having to handle multiple addresses when sending requests to a docker server

Examples

```
http://machine1:8080{some constant path for service 1}
http://machine1:8081{some constant path for service 2}
http://machine1:8082{some constant path for service 3}
```

Would now be

```
http://machine1:8090{some constant path for service 1}
http://machine1:8090{some constant path for service 2}
http://machine1:8090{some constant path for service 3}
```

Depending how your configuration is done it effectively turns

```
var service1Url = {service 1 base address}{some constant path for service 1}
var service2Url = {service 2 base address}{some constant path for service 2}
var service3Url = {service 3 base address}{some constant path for service 3}
```

Into

```
var service1Url = {service api base address}{some constant path for service 1}
var service2Url = {service api base address}{some constant path for service 2}
var service3Url = {service api base address}{some constant path for service 3}
```

## How to use

``` yml
version: "2"
services:
  # This is the proxy container and is the only one that needs to have port mapping unless
  # it is being accessed through another container proxy
  proxy:
    image: bspjojo/docker-proxy:latest
    container_name: proxy
    ports:
     - "8090:80"
    environment:
      - VIRTUAL_HOST=whoami.localhost
    volumes: 
      - /var/run/docker.sock:/var/run/docker.sock:ro
  # This is a container that is ignored by the proxy container
  web:
    image: nginx
    container_name: "web"
    ports:
      - "8080:80"
    environment:
      - NGINX_HOST=foobar.com
      - NGINX_PORT=80
  # The proxy container will set up a proxy_pass for this container
  # All requests at /api/service1 will go here
  api1:
    image: bspjojo/testing-api:latest
    environment:
      - PROXY_LOCATION=/api/service1
  # The proxy container will set up a proxy_pass for this container
  # All requests at /api/service2 will go here
  api2:
    image: bspjojo/testing-api:latest
    environment:
      - PROXY_LOCATION=/api/service2
  # The proxy container will set up a proxy_pass for this container
  # All requests at / or /app will go here
  extra:
    image: nginx
    container_name: "extra"
    environment:
      - PROXY_LOCATION=/,/app
    volumes: 
      - ./testing-page/index.html:/usr/share/nginx/html/index.html
      - ./testing-page/index.html:/usr/share/nginx/html/app/index.html
  # https://github.com/nginx-proxy/nginx-proxy
  # Used to proxy based on hostname in the request
  # Not needed when accessing the proxy container directly
  nginx-proxy:
    image: jwilder/nginx-proxy
    container_name: nginx-proxy
    ports:
      - "80:80"
    volumes:
      - /var/run/docker.sock:/tmp/docker.sock:ro
```

The proxy container will use the exposed port on any container with an environment variable named `PROXY_LOCATION`.

The proxy container must have access to the docker service.

If you need to support multiple locations on a single container, you can separate each location with commas. For example, `PROXY_LOCATION=/ap1/service3,/ap1/service4`. Requests to both `/api/service3` and `/api/service4` will both go to that same container.

## Limitations

Currently only supports

- Having one exposed port on any container
- HTTP between containers
- Containers must be on the same network as the proxy container
- Containers must have `PROXY_LOCATION` environment variable
- Linux containers only
