# DESIGN.md

# Event Tracking System – Design Overview

### Architecture

The solution follows a layered architecture to separate responsibilities and maximize maintainability.

```text
React/Remix Frontend
        │
        ▼
Gin HTTP API
        │
        ▼
Authentication / Request Enrichment
        │
        ▼
Generic Pipeline Library (Pipeline[T])
        │
        ▼
Batch Runtime & Worker Pool
        │
        ▼
Repository Layer
        │
        ▼
ClickHouse
```

The frontend captures user interactions using standard Web APIs and sends batched tracking events to the Go API. The API authenticates requests using JWT stored in an HttpOnly cookie, enriches events with server-side request metadata, processes them through a reusable generic pipeline library, and persists both events and pipeline execution metrics to ClickHouse. Prometheus metrics are exposed for operational monitoring.

---

## Key Design Decisions

### 1. Generic Pipeline Library (`Pipeline[T]`)

### Decision

The pipeline library uses a homogeneous generic type (`Pipeline[T]`) instead of type-erased (`any`) stage boundaries.

### Why

This provides compile-time type safety, eliminates runtime casting, improves IDE support, and reduces runtime overhead. Since the pipeline processes a single event type throughout execution, a homogeneous model naturally fits the problem.

### Alternative Considered

Using `any` between stages would allow heterogeneous pipelines but would introduce runtime type assertions, increase implementation complexity, and weaken compile-time guarantees. Given the assignment scope, the additional flexibility was not justified.

---

### 2. Concurrent Runtime with Dynamic Batching

### Decision

The runtime uses a configurable worker pool, buffered queues, and dynamic batching.

### Why

Batching significantly reduces ClickHouse insert overhead, while multiple workers improve throughput on multi-core systems. Worker count, batch size, flush interval, and channel depth are configurable to support different workloads without code changes.

### Alternative Considered

Processing every event synchronously was simpler but would not scale under sustained load due to excessive database calls and poor CPU utilization.

---

### 3. JWT Authentication Using HttpOnly Cookies

### Decision

Authentication uses JWT stored in an HttpOnly cookie.

### Why

This satisfies the assignment requirements while reducing exposure to XSS attacks compared to storing tokens in browser storage. The browser automatically includes the cookie on authenticated requests, simplifying client-side authentication.

### Trade-off

Cross-origin deployments require careful CORS and cookie configuration (`SameSite=None` and `Secure=true`). For local development, a proxy was used to maintain a same-origin experience.

---

## 4. ClickHouse for Analytics Storage

### Decision

ClickHouse is used as the analytics datastore for both tracking events and pipeline execution metrics.

### Why

ClickHouse is optimized for high-volume analytical workloads, supports efficient batched inserts, and enables fast aggregation queries for event statistics, latency, throughput, and operational dashboards.

### Alternative Considered

A transactional database such as PostgreSQL was considered but rejected because the workload is append-heavy and analytical rather than transactional.

# Trade-offs

### Pipeline Typing

The homogeneous generic design offers strong compile-time safety and simpler implementations but does not allow pipelines to transform between unrelated types without explicit mapping. This trade-off favors maintainability and performance over maximum flexibility.

### Authentication Model

Using HttpOnly cookies improves security and aligns with browser-based authentication but requires additional CORS configuration when frontend and backend are hosted on different origins.

### Dynamic Batching

Larger batches improve throughput and reduce database overhead but increase latency for individual events. The runtime therefore exposes batch size and flush interval as configurable parameters to balance throughput and responsiveness.

## Scaling Strategy

The current implementation is suitable for a single service instance and can evolve without architectural changes.

Potential scaling improvements include:

* Horizontally scale multiple API instances behind a load balancer.
* Introduce Kafka or another durable message broker between the API and analytics pipeline.
* Use ClickHouse replication and sharding for increased ingestion capacity.
* Auto-tune worker count and batch size based on runtime queue depth and latency.
* Introduce dead-letter queues for permanently failed event batches.

## Team Work Plan

For Three Engineers, I would divide the work across three engineers while maintaining well-defined integration points.

**Engineer A – Frontend & Authentication**

* Build the tracking SDK using Standard Web APIs.
* Implement sign-up, login, and JWT cookie flow.
* Capture user interactions and batch events.

**Engineer B – Pipeline & Runtime**

* Develop the generic pipeline library.
* Implement worker pool, batching, runtime, metrics, benchmarks, and tests.

**Engineer C – Backend & Analytics**

* Implement the Gin API, ClickHouse repositories, event ingestion, Prometheus integration, dashboards, and deployment.

As Technical Lead, I would own the overall architecture, define interfaces between components, review critical design decisions, coordinate integration, and monitor delivery risks.

**Phase 1:** Finalize architecture, API contracts, database schema, authentication, pipeline interfaces, and basic end-to-end integration.

**Phase 2:** Complete pipeline library with concurrent runtime, metrics, dashboards, testing, benchmarking, documentation, and final review.
