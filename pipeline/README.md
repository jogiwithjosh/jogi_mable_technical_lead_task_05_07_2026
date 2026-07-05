

# Pipeline Design Decisions
## Typing Model
### Chosen Approach: Homogeneous Generic Pipeline (Pipeline[T])

The pipeline library is implemented using Go generics:

```
type Pipeline[T any] struct {
    ...
}
```
Each stage operates on the same element type T.

Examples:
```
pipe := pipeline.New[event.Event](cfg).
    Map("validate", validateEvent).
    Map("normalize", normalizeEvent).
    Filter("remove-invalid", isValidEvent).
    Generate("derived-events", createDerivedEvents)
```
Every stage accepts and emits event.Event.

### Why this approach? 
#### Advantages
- Compile-time type safety - no runtime casting is required.
    ```
    func(ctx context.Context, e event.Event) (event.Event, error)

        instead of

    func(any) any

        This eliminates an entire class of runtime errors.
    ```

- Higher performance

    Because the implementation is generic, Go generates specialized code without repeated interface conversions.

    There is no need for: `event := value.(Event)` inside every stage.

    This reduces runtime overhead in a high-throughput ingestion pipeline.

#### Trade-offs

The main limitation is that every stage must use the same element type.

This means a stage cannot naturally transform:
```
Event -> User or Event -> AnalyticsRecord
```

without wrapping the additional data inside the existing type.

For this analytics ingestion system, all processing stages operate on a single event model, so a homogeneous pipeline is a natural fit.

If future requirements involve heterogeneous transformations, the pipeline could be extended with typed stage interfaces or separate pipelines connected through adapters.

## Why not any?

```
func(any) any
```

#### Advantages:

maximum flexibility
heterogeneous stage types
arbitrary transformations

#### Disadvantages:

- runtime type assertions
- reduced performance
- loss of compile-time safety
- harder debugging

this library targets high-throughput processing, compile time safety and performance were prioritised over maximum flexibility.

## Benchmark
```
MacBook M1 Pro
8 CPU cores
16 GB RAM
GOMAXPROCS=10

goos: darwin
goarch: arm64
pkg: pipeline
cpu: Apple M1 Pro

BenchmarkPipeline10-10                    824019              1473 ns/op            3504 B/op         14 allocs/op
BenchmarkPipeline1K-10                     12992             88920 ns/op          303265 B/op       1004 allocs/op
BenchmarkPipeline100K-10                     147           8135100 ns/op        29603424 B/op     100004 allocs/op
BenchmarkPipeline1M-10                        18          63000273 ns/op        296006496 B/op   1000004 allocs/op
BenchmarkRuntime1M_4W500B-10                   5         203986717 ns/op        196834217 B/op      6001 allocs/op
BenchmarkRuntime10M_4W500B-10                  1        2763431041 ns/op        1968324992 B/op    60012 allocs/op
BenchmarkRuntime1M_1W128B-10                   5         221398600 ns/op        212874899 B/op     23438 allocs/op
BenchmarkRuntime1M_8W500B-10                   6         210492035 ns/op        196832242 B/op      6002 allocs/op
BenchmarkRuntime1M_8W1024B-10                  6         181703972 ns/op        192110808 B/op      2930 allocs/op
BenchmarkRuntime1M_16W500B-10                  6         201222674 ns/op        196833829 B/op      6000 allocs/op
```

## Benchmark Summary
| Workers | Batch Size |         Time |          Throughput |       Memory |    Allocs |
| ------: | ---------: | -----------: | ------------------: | -----------: | --------: |
|       4 |        500 |     201.5 ms |     4.96 M events/s |     196.8 MB |     6,000 |
|       8 |        500 |     221.9 ms |     4.51 M events/s |     196.8 MB |     6,000 |
|       8 |       1024 |     304.6 ms |     3.28 M events/s |     192.1 MB |     2,931 |
|      16 |        500 |     326.8 ms |     3.06 M events/s |     196.8 MB |     6,002 |


## Supported Stages

The library supports the following stage types:

* **Map** – Transform an event.
* **Filter** – Drop events based on a predicate.
* **Generate** – Produce zero or more derived events from a single input event.
* **If** – Route events into one of two sub-pipelines and merge the outputs.
* **Reduce** – Aggregate a stream into a single result (terminal stage).
* **Sink** – Drain processed events into a bounded sink.

New stage types can be added by implementing the `Stage[T]` interface without modifying the pipeline core.

---

## Extensibility

The pipeline is designed to be open for extension and closed for modification. New stage types can be added by implementing the common Stage[T] interface without changing the pipeline core.

```
Stage Interface

type Stage[T any] interface {
    Name() string
    Process(ctx context.Context, items []T) ([]T, error)
}
```
Every pipeline stage (Map, Filter, Generate, If, Reduce, Collect) implements this interface.

To add a new stage:

Create a new struct implementing Stage[T].
Implement the Process() method.
Register/add the stage to the pipeline.

No changes are required in the pipeline runtime or executor.

---

## Runtime Architecture

Incoming events are processed asynchronously using a configurable runtime.

```text
HTTP API
    │
    ▼
Buffered Queue
    │
    ▼
Dynamic Batcher
    │
    ▼
Worker Pool
    │
    ▼
Pipeline Executor
    │
    ▼
 Sink
```

The runtime supports:

* Concurrent worker pool
* Dynamic batching
* Buffered queues
* Graceful shutdown
* Backpressure
* Metrics collection

---

## Configurable Parameters

```go
Config{
    Workers:    4,
    BatchSize:  500,
    BufferSize: 10000,
    FlushEvery: time.Second,
}
```

| Parameter      | Purpose                                           |
| -------------- | ------------------------------------------------- |
| **Workers**    | Number of concurrent batch processing goroutines  |
| **BatchSize**  | Maximum events processed together                 |
| **BufferSize** | Queue capacity between producers and workers      |
| **FlushEvery** | Maximum wait time before flushing a partial batch |

These parameters allow the pipeline to be tuned for different workloads.

---

## Performance Tuning

Several configurations were benchmarked to determine the optimal defaults.

| Workers | Batch Size |   1M Events |
| ------: | ---------: | ----------: |
|       1 |        128 |     ~340 ms |
|   **4** |    **500** | **~202 ms** |
|       8 |        500 |     ~222 ms |
|      16 |        500 |     ~327 ms |

Based on these results, the default configuration of **4 workers** and **batch size 500** provided the best balance of throughput, memory usage, and synchronization overhead on the benchmark machine.

---

## Pipeline Metrics

The runtime collects per-stage metrics including:

* Processing latency
* Batch size
* Throughput
* Dropped events
* Error count
* Queue depth

These metrics are exposed through Prometheus and can be stored alongside tracking events for operational monitoring and performance analysis.

