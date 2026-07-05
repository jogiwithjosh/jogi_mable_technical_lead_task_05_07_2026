                              Browser
                                 │
                                 │
                 ┌───────────────┴────────────────┐
                 │                                │
                 │      React 19 + Remix          │
                 │                                │
                 │ Authentication                 │
                 │ Product Library                │
                 │ Cart                           │
                 │ Checkout                       │
                 │ Tracking Script/SDK                   │
                 │ Error Boundary                 │
                 └───────────────┬────────────────┘
                                 │
                    HttpOnly JWT │
                                 │
                         REST API (Gin)
         Authentication                 Event API              Product API                        │ 
       ┌─────────────────────────┼────────────────────────┐
       │                         │                        │
       │                         │                        │
       |                         |                        |
       │                                                  │
       │                  Pipeline Library                │
       │                         │                        │
       │                  Validation Stage                │
       │                  Normalization                   │
       │                  Enrichment                      │
       │                  Metrics                         │
       │                         │                        │
       └─────────────────────────┴────────────────────────┘
                                 │
                          ClickHouse DB
                                 │
                           Grafana Dashboard


## Tracking Script / SDK
### Tech Stack

#### Idea is to make trackng script frontend framework agnostic, making it reusable across different frameworks/applications.

- TypeScript - strong compile-time guarantees and safer refactoring.
- Vite - Vite outputs both ESM and CommonJS bundles with TypeScript declarations.

- No React dependency
- No analytics libraries
- No external state libraries

## ECommerce App

## Go API
~~~
                React E-commerce
                        │
                        │
        Tracking SDK (Batch)
                        │
                        ▼
                Gin REST API
                        │
          Authentication Middleware
                        │
                        ▼
            pipeline.Pipeline[Event]
                        │
        ┌───────────────┼────────────────┐
        ▼               ▼                ▼
   Validation      Normalization    Enrichment
                        │
                        ▼
              Analytics Service
                        │
                        ▼
           ClickHouse Repository
                        │
                        ▼
                  ClickHouse DB
                        │
              ┌─────────┴─────────┐
              ▼                   ▼
          Statistics          Grafana
~~~

~~~
                    React SPA
                        │
                        │
                 http://localhost:5173
                        │
                        ▼
                 Gin API :8080
                        │
          ┌─────────────┴────────────┐
          ▼                          ▼
     ClickHouse                 Prometheus
          │                          │
          └─────────────┬────────────┘
                        ▼
                    Grafana
~~~

## pipeline go SDK
~~~
Event
   │
   ▼
Pipeline.Execute()

        │
        ▼

┌──────────────────────┐
│ Validation Stage     │
├──────────────────────┤
│ Normalization Stage  │
├──────────────────────┤
│ Enrichment Stage     │
└──────────────────────┘

        │
        ▼

Processed Event
~~~