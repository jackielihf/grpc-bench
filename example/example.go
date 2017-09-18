
package main

import(
    "fmt"
    "flag"
    "time"
    "math/rand"
    bench "github.com/jackielihf/grpc-bench"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "github.com/jackielihf/grpc-bench/api/grpcapi"
)

var addr *string = flag.String("grpc", "localhost:5000", "grpc address")
var total *int = flag.Int("n", 1, "total number of requests")
var c *int = flag.Int("c", 1, "number of concurrency/client")
var p *int = flag.Int("p", 1, "size of pool of client")


func handler(conn *grpc.ClientConn) error {

    api := grpcapi.NewAccountManagerClient(conn)    
    req := new(grpcapi.QueryProjectReq)
    pageInfo := new(grpcapi.PageInfo)
    pageInfo.Page = 0
    pageInfo.Limit = 10
    req.PageInfo = pageInfo
    req.ClientId = 151691785431679090
    _, err := api.ListProjectByClientId(context.Background(), req)

    return err
}

func handler2(conn *grpc.ClientConn) error {

    // do something
    // t := int(time.Now().UnixNano() % 10)
    fmt.Println("run ...")
    time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)

    return nil
}

func main() {

    flag.Parse()

    bench.Benchmark(*addr, *total, *c, *p, handler)

}