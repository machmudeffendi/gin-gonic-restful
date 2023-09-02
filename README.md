# Gin Gonic Restful API

# Prerequisites
- golang v1.21^
- PostgreSQL

# Instalation
```bash
$ make init-dependency
```

# Run Dev
```
$ go run main.go
```

# Build 
will reduce compiled size from 20MB to 15MB 
```
$ go build -ldflags "-s -w"
```
Run with
```
$ ./restfull-gin-gonic
```