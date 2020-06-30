# C100K MVP

## Technology Stack

- MySQL
- Redis
- Go
- Docker


## API

- User Auth
- Goods List
- Goods Detail
- Goods Search
- Cart
- Create Order
- Pay Order


## Benchmarks

wrk -c36 -t12 -d 5s http://127.0.0.1:8018

```
Running 5s test @ http://127.0.0.1:8018
  12 threads and 36 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   347.50us  595.77us  37.92ms   96.23%
    Req/Sec    10.16k     2.96k   41.40k    78.07%
  608907 requests in 5.10s, 78.39MB read
Requests/sec: 119396.74
Transfer/sec:     15.37MB
```
