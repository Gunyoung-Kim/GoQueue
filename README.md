# Queue in Go

### Purpose

Start this project to use Queue data structure which is not in Go standard library. 


## Getting Started

### Installing

Install Go and run `go get`:

```sh
$ go get github.com/Gunyoung-Kim/GoQueue/...
```

### Testing 

```sh
$ go test ./...
```

### Performance

```sh 
$ go test ./... -bench=.
goos: darwin
goarch: amd64
pkg: github.com/Gunyoung-Kim/queue/queue/queue
cpu: Intel(R) Core(TM) i5-8279U CPU @ 2.40GHz
BenchmarkAdd-8             55870             20906 ns/op
BenchmarkAddByList-8       23574             51621 ns/op
PASS
ok      github.com/Gunyoung-Kim/queue/queue/queue       3.453s
```
Compare with container/list in Go standard library

### Create Queue 


```go
package main

import "github.com/Gunyoung-Kim/GoQueue/queue"

func main() {
	// create new empty Queue
	newQueue := queue.NewQueue()
	// add new element in queue
	newQueue.Add(1)
}

```

---

Kim GunYoung
