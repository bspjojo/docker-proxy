# Docker proxy

## What is this?

A simple proxy container that proxies requests through to other containers

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
  # This is the proxy container and is the only one that needs to have port mapping
  proxy:
    image: bspjojo/docker-proxy:latest
    container_name: proxy
    ports:
     - "8090:80"
    volumes: 
      - /var/run/docker.sock:/var/run/docker.sock:rw
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
  # This is a container that is ignored by the proxy container
  extra:
    image: nginx
    container_name: "extra"
    ports:
    - "8081:80"
    environment:
    - NGINX_HOST=foobar.com
    - NGINX_PORT=80
```

The proxy container will use the exposed port on any container with an environment variable named `PROXY_LOCATION`
The proxy container must have access to the docker service

## Limitations

Currently only supports

- Having one exposed port on any container
- One proxy location
- HTTP between containers
- Containers must be on the same network as the proxy container
- Containers must have `PROXY_LOCATION` environment variable
- Linux containers only
