# Project 

This project's intention is to try the performance of Go-based application backed with Redis as a data source against load testing

## Technology

- API framework: `Gin`
- Datasource: `Redis`
- Load testing tool: [k6](https://k6.io/) 

## How to test

run `make run` to start Docker compose which will start Redis and your service
run `k6 run -d 30s -u 30 ./load-test-script.js` to start testing with
- 30 VUs (virtual users)
- and for Duration of 30 seconds