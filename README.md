# gRPC Performance Comparison: Go vs Rust

This Pet project compare the performance of gRPC implementations in Go and Rust (using Tonic). 
By running performance tests, we can analyze and understand the efficiency and speed of gRPC servers written in these two languages.

> [!NOTE]  
> I don't want to start a language war. My goal is simply to improve my understanding to make better decisions in terms of technology.

![img](docs/results.png)

## Requirements

- Go (latest version)
- Rust (latest version)
- `ghz` (gRPC benchmarking and load testing tool `brew install ghz`)

## Build Instructions

To build the project, simply run the provided build script:

```sh
./build.sh
```

## Run it Yourself

### Start the Server to Test

You can start the gRPC server in Go or Rust by running the appropriate binary:

For Go:
```sh
bin/go_server
```

For Rust:
```sh
bin/rust_server
```

### Load Test the Server

Once the server is running, you can perform a load test using `ghz`:

```sh
ghz --config=ghz_config.json --total 1000000  --concurrency 5000 --connections 100 -e
```

This command will execute a performance test with 1 Million total requests, 5000 concurrent requests, and 100 connections.


## Results

# Go Results
```
ghz --config=ghz_config.json --total 1000000  --concurrency 5000 --connections 100 -e

Summary:
  Count:        1000000
  Total:        12.92 s
  Slowest:      578.58 ms
  Fastest:      0.04 ms
  Average:      41.43 ms
  Requests/sec: 77424.85

Response time histogram:
  0.035   [1]      |
  57.890  [772292] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  115.744 [120096] |∎∎∎∎∎∎
  173.599 [64960]  |∎∎∎
  231.453 [26749]  |∎
  289.308 [10820]  |∎
  347.162 [3682]   |
  405.017 [1118]   |
  462.871 [131]    |
  520.726 [93]     |
  578.580 [58]     |

Latency distribution:
  10 % in 1.25 ms 
  25 % in 5.85 ms 
  50 % in 17.92 ms 
  75 % in 51.28 ms 
  90 % in 120.40 ms 
  95 % in 164.72 ms 
  99 % in 258.08 ms 

Status code distribution:
  [OK]   1000000 responses   
```

* Note: results on a Apple M3 Max with 53 total Memory usage and 200% cpu usage.


# Rust results

```
ghz --config=ghz_config.json --total 1000000  --concurrency 5000 --connections 100 -e

Summary:
  Count:        1000000
  Total:        13.35 s
  Slowest:      728.93 ms
  Fastest:      0.03 ms
  Average:      44.69 ms
  Requests/sec: 74881.26

Response time histogram:
  0.033   [1]      |
  72.923  [789859] |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  145.813 [130969] |∎∎∎∎∎∎∎
  218.703 [52996]  |∎∎∎
  291.593 [19038]  |∎
  364.483 [5744]   |
  437.373 [996]    |
  510.263 [310]    |
  583.153 [68]     |
  656.043 [5]      |
  728.933 [14]     |

Latency distribution:
  10 % in 1.22 ms 
  25 % in 5.37 ms 
  50 % in 18.38 ms 
  75 % in 57.00 ms 
  90 % in 131.86 ms 
  95 % in 175.49 ms 
  99 % in 275.52 ms 

Status code distribution:
  [OK]   1000000 responses   

```
* Note: results on a Apple M3 Max with 53 total Memory usage and 160% cpu usage.

## Result Overview

In my small quest to compare gRPC performance between Go and Rust, I found that both languages deliver impressive results, with Go edging out Rust with slightly better performance metrics!

I believe that with some optimizations, Rust, being a non-GC language, could potentially outperform Go. 

However, given the impressive performance numbers we've seen and all the other aspects to consider when building a new service, there are only a few cases where Rust would make a significant difference. 

In my personal opinion, the added complexity of Rust isn't necessary if you're looking for a high-performance gRPC service. All things considered, it's mind-blowing how effortlessly you can write massive concurrent applications nowadays!

