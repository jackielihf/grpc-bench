
#!/bin/bash
cd $(cd `dirname $0`;pwd)

CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o ../bin/bench_linux ../example/example.go
CGO_ENABLE=0 GOOS=darwin go build -a -installsuffix cgo -o ../bin/bench_darwin ../example/example.go
