# Bench Srv

This is the Bench service with fqdn go.micro.srv.bench.

## Getting Started

### Prerequisites

Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -dev -advertise=127.0.0.1
```

DB table
```
CREATE TABLE test ( id INT NOT NULL AUTO_INCREMENT, foo VARCHAR(100), bar VARCHAR(100), PRIMARY KEY(id));
```

### Run Service

```
$ go run main.go
```

### Building a container

If you would like to build the docker container do the following
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o bench-srv ./main.go
docker build -t bench-srv .

```
