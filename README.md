# dummy-api
this api is used for testing purposes

To run:
```bash
go run .
```

test:
```bash
#health check
curl localhost:7000/health

#read
curl localhost:7000/read?latency=1000

#write
curl -X POST localhost:7000/write?latency=1000
```
