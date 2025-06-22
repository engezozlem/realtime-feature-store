# Realtime Feature Store

A polyglot, real-time feature store system designed for ML systems that require low-latency, consistent feature access
across online and offline contexts.

## Overview

This project demonstrates how to build a production-grade feature store combining:

- **Go** → Fast, stateless HTTP API server
- **Redis** → Low-latency key-value feature storage
- **Python** → Model training, feature logging, MLflow integration (next step)
- **Spark / Flink** → Batch and streaming ETL jobs (future phase)

---

## Components

| Folder      | Language | Description                                        |
|-------------|----------|----------------------------------------------------|
| `api/`      | Go       | HTTP endpoint handlers                             |
| `store/`    | Go       | Redis client and feature data access layer         |
| `infra/`    | Docker   | Local dev environment (`docker-compose.yml`)       |
| `ml/`       | Python   | ML notebooks and tracking setup (MLflow)           |
| `ingestor/` | Py/Scala | Future batch/streaming feature ingestion code      |
| `docs/`     | Markdown | Architecture information and project documentation |
| `README.md` | —        | This file                                          |

---

## Getting Started

### 1. Clone and enter project

```bash
git clone git@github.com:engezozlem/realtime-feature-store.git
cd realtime-feature-store
```

### 2. Start Redis locally

``` bash
docker-compose -f infra/docker-compose.yml up -d
```

### 3. Run the Go API server

``` bash
go mod tidy
go run main.go
```

### 4. Insert sample feature data

``` bash
docker ps # find the Redis container name (usually infra-redis-1)
docker exec -it infra-redis-1 redis-cli
HSET 123 user_age 42 country "TR"
```

### 5. Test the feature API

``` http request
curl http://localhost:8080/features/123
```

You should see:

``` json
{
  "entity_id": "123",
  "features": {
    "user_age": "42",
    "country": "TR"
  }
}
```

# Roadmap

1. [X] Go API + Redis integration
2. [X] Basic feature retrieval with fallback
3. [ ] ML pipeline + MLflow integration (ml/)
4. [ ] Spark / Flink streaming ingestion pipeline (ingestor/)
5. [ ] Feature consistency validator
6. [ ] Drift detection tooling
7. [ ] Optional LLM-assisted retraining triggers

