# 🚀 Realtime Feature Store

A cross-language, real-time machine learning feature store using **Python**, **Go**, **Redis**, and **MLflow** — built for learning, experimenting, and scaling real-time ML systems from scratch.

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/engezozlem/realtime-feature-store.git
cd realtime-feature-store
```

---

### 2. Start Docker Services

Redis & MLflow UI:

```bash
docker-compose up -d
```

- Redis: `localhost:6379`
- MLflow UI: `http://localhost:5000`

---

### 3. Set Up Python Environment (for ingestion)

```bash
cd ml
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
python ingestor.py
```

> This writes feature data (from `sample_users.csv`) into Redis using `HSET`.

---

### 4. Start Go Feature API

```bash
cd ../
go run main.go
```

Then test it:

```bash
curl http://localhost:8080/features/123
```

Expected output:

```json
{
  "entity_id": "123",
  "features": {
    "user_age": "42",
    "country": "TR"
  }
}
```

## 🔧 Technologies

- Python 3.10
- Go 1.21+
- Redis 6/7
- MLflow 2.x
- Docker

---

## 🧠 What’s Next?

- TTL support for feature freshness
- Streaming ingestion (Kafka/Flink or simulated)
- Real-time model serving (FastAPI or BentoML)
- Observability with Prometheus + Grafana
- CI setup with GitHub Actions

---