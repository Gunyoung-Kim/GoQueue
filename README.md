# Queue in Go

### Purpose

Start this project to use Queue data structure which is not in Go standard library. 


## Getting Started

### Installing

Install Go and run `go get`:

```sh
$ go get github.com/Gunyoung-Kim/GoQueue/...
```

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

Kim GunYoung (Korean)
