# dummy-api
this api is used for testing purposes

To run:
```bash
go run .

#with docker
docker rm dummy-api; docker run -p 7000:7000 --name dummy-api -it ghcr.io/elhmn/dummy-api:d5f568d3b3ed9bd5cf20fffe53ac96777d6ad72d
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
