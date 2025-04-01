# Reverse Proxy

A simple reverse proxy application. The initial code comes from
https://gist.github.com/JalfResi/6287706.

## Usage

Usage: reverse-proxy [OPTIONS]

    Options:
      -r, --remote <REMOTE>
      -l, --local <LOCAL>

## Example

Start a reverse proxy on localhost:3030 forwarding to localhost:8080.
```
reverse-proxy --local 0.0.0.0:3030 --remote http://localhost:8080
```

Start a tcp server on localhost bind to tcp port 8080 using netcat
```
# netcat -l 0.0.0.0 8080
```

Use curl to perform an HTTP request
```
# curl -v "http://127.0.0.1:3030/hello/my/friend"
> GET /hello/my/friend HTTP/1.1
> Host: 127.0.0.1:3030
> User-Agent: curl/8.8.0
> Accept: */*
```

Query received by netcat server
```
GET / HTTP/1.1
host: 127.0.0.1:3030
accept: */*
```
