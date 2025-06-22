# Architecture: Realtime Feature Store

This document explains the design and rationale behind the realtime feature store, covering key components, serving
strategies, system flow, and future extensibility.

---

## Overview

This system enables **low-latency, online feature retrieval** for ML systems, ensuring consistency between training and
inference time through modular, polyglot components.

Core technologies:

- **Go** — High-performance API server
- **Redis** — Low-latency key-value store for real-time feature access
- **Docker Compose** — Simplified development environment
- *(Future: Spark/Flink for ETL, MLflow for tracking, Python/Scala for modeling)*

---

## Serving Strategy

### Pull-based Serving (Current)

- API server retrieves features from Redis at inference time
- Supports dynamic, on-demand entity queries
- Redis acts as a real-time feature cache

**Pros:**

- Always up-to-date
- Stateless API design
- Works well with user-facing prediction endpoints

**Cons:**

- Requires Redis availability
- Slight latency cost on each request

---

## Push-based Serving (Future)

- Precomputed features pushed to Redis periodically (via Flink/Spark jobs)
- Works for batch predictions or streaming ML pipelines

**Pros:**

- Low latency once materialized
- Simplifies edge device predictions

**Cons:**

- Requires ETL orchestration
- Risk of staleness

---

## 🧱 System Components

```text
         ┌────────────┐
         │   Client   │
         └─────┬──────┘
               │ HTTP
        ┌──────▼───────┐
        │   Go API     │
        │ /features/:id│
        └──────┬───────┘
               │ Redis GET
        ┌──────▼───────┐
        │   Redis      │
        └──────────────┘
```

## Entity-Feature Model

Entities are identified by an entity_id. Feature values are stored in Redis using HSET:

```bash
HSET 123 user_age 42 country "TR"
```

Features are retrieved via:

```http request
GET /features/123
```

Returns:

```json
{
  "entity_id": "123",
  "features": {
    "user_age": "42",
    "country": "TR"
  }
}
```

# Next Architectural Milestones

* 🟡 Stream ingestion layer (Flink/Spark → Redis)
* 🟡 MLflow integration for training-time logging
* 🟡 Snapshot comparator for online/offline parity
* 🟡 Feature expiration with Redis TTL
* 🟢 Model agent for retraining triggers (LLM-assisted, optional)

# Design Principles

* Polyglot-first: Go, Python, Scala coexist cleanly
* Pluggable architecture: Easily extend components or replace Redis
* Observability-ready: Logging, tracking, validation in mind
* Modular codebase: API, store, ML, and infra clearly separated