
package bench

import(
    "sync"
    "sync/atomic"
    "fmt"
    "time"
    "log"
    "github.com/montanaflynn/stats"
    "google.golang.org/grpc"
)


type Handler func(*grpc.ClientConn) error

func Benchmark(grpcAddr string, total int, concurrency int, pool int, handler Handler) {

    c := concurrency
    if c < 1 {
        c = 1
    }
    if total < 1 {
        total = 1
    }
    if pool < 1 {
        pool = 1
    }
    m := int(total / c) // request times of one thread
    var wg sync.WaitGroup

    fmt.Printf("\nBenchmark running:\ngrpc-address: %s, concurrency: %d, pool: %d, expected total request: %d\n\n", grpcAddr, c, pool, total)

    // data
    var reqs uint64 = 0
    var doneReqs uint64 = 0

    tArray := make([]float64, c * m)

    startTime := time.Now()



    for i := 0; i < c; i++ {

        var sem = make(chan int, pool)
        wg.Add(1)

        go func(index int) {

            var cwg sync.WaitGroup

            // connect
            conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
            if err != nil {
                log.Fatalf("grpc dial failed: err=%v", err)
            }

            defer func () {
                // conn.Close()
            }()


            // run
            for j := 0; j < m; j++ {
                t := time.Now()
                sem <- 1 // wait for resource
                cwg.Add(1)

                go func(offset int) {

                    err := handler(conn)    

                    // record
                    if err == nil {
                        atomic.AddUint64(&doneReqs, 1)
                    }
                    tArray[index * m + offset] = float64((time.Now().UnixNano() - t.UnixNano()) / 1000000)
                    atomic.AddUint64(&reqs, 1)

                    // release resource
                    <-sem
                    cwg.Done()
                }(j)
            }
            cwg.Wait()
            wg.Done()

        }(i)
    }

    wg.Wait()
    // stat
    totalTime := (time.Now().UnixNano() - startTime.UnixNano()) / 1000000
    mean, _ := stats.Mean(tArray)
    max, _ := stats.Max(tArray)
    min, _ := stats.Min(tArray)

    if totalTime < 1 {
        totalTime = 1
    }
    rqs := int64(c * m * 1000) / totalTime

    // print
    fmt.Printf("--------------------\n")
    fmt.Printf("Total requests: %d\n", reqs)
    fmt.Printf("Successful requests: %d\n", doneReqs)
    fmt.Printf("Total elapse time: %dms\n", totalTime)

    fmt.Printf("Latency (mean, min, max): %.fms  %.fms  %.fms\n", mean, min, max)
    fmt.Printf("Average Throughput: %d\n", rqs)
    fmt.Printf("\n\n")
}


