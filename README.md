# API Benchmark

This repository contains two CLI tools written in Go:

- **API Benchmark Tool**: Perform load testing on APIs or websites by sending concurrent HTTP requests.

---

## Installation

Make sure you have Go installed (version 1.16+).

Clone the repo and run the tool using:

```bash
go run main.go [command] [flags]
```


## Usage

### Benchmark Tool (default command)
```
go run main.go --url https://httpbin.org/get --requests 100 --concurrency 10
```

### Flags:

| Flag          | Shortcut | Description                  | Default |
| ------------- | -------- | ---------------------------- | ------- |
| --url         | -u       | Target URL to benchmark      | (none)  |
| --requests    | -r       | Total number of requests     | 100     |
| --concurrency | -c       | Number of concurrent workers | 10      |


### Example output:
```
Benchmark Results:
Total Requests:       100
Concurrency:          10
Total Time:           5.03s
Requests per Second:  19.87
Average Response Time:503.8ms
Min Response Time:    150ms
Max Response Time:    1.2s
P95 Response Time:    900ms
P99 Response Time:    1.1s
Failed Requests:      0
Status Code Summary:
  200 => 100 responses
```