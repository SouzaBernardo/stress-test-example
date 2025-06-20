kotlin locally

Running 30s test @ http://localhost:8080/hello
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.25ms    6.76ms 194.19ms   98.68%
    Req/Sec    15.21k     3.29k   21.86k    72.91%
  1810432 requests in 30.01s, 138.12MB read
Requests/sec:  60335.64
Transfer/sec:      4.60MB
bepssouza@Bernardos-MacBook-Pro stress-test-example % wrk -t4 -c100 -d30s http://localhost:8080/hello
Running 30s test @ http://localhost:8080/hello
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.63ms    1.35ms  33.89ms   90.12%
    Req/Sec    16.95k     1.84k   21.50k    68.08%
  2023813 requests in 30.00s, 154.40MB read
Requests/sec:  67451.42
Transfer/sec:      5.15MB
bepssouza@Bernardos-MacBook-Pro stress-test-example % wrk -t4 -c100 -d30s http://localhost:8080/hello
Running 30s test @ http://localhost:8080/hello
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.08ms    7.69ms 184.35ms   99.08%
    Req/Sec    17.27k     2.14k   21.69k    72.07%
  2059035 requests in 30.01s, 157.09MB read
Requests/sec:  68613.15
Transfer/sec:      5.23MB